package handlers

import (
	"avito-api/internal/avito-api/models"
	"avito-api/internal/avito-api/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type HouseHandler struct {
	HouseService *services.HouseService
	FlatService  *services.FlatService
}

func NewHouseHandler(houseService *services.HouseService, flatService *services.FlatService) *HouseHandler {
	return &HouseHandler{
		HouseService: houseService,
		FlatService:  flatService,
	}
}

func (h *HouseHandler) CreateHouse(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод недоступен", http.StatusMethodNotAllowed)
		return
	}

	var houseInputObject models.HouseInputObject
	if err := json.NewDecoder(r.Body).Decode(&houseInputObject); err != nil || houseInputObject.Address == "" || houseInputObject.Year <= 0 {
		http.Error(w, "Невалидные данные ввода", http.StatusBadRequest)
		return
	}

	house, err := h.HouseService.CreateHouse(&houseInputObject)
	if err != nil {
		log.Printf("Error inserting house: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(house)
}

func (h *HouseHandler) GetFlatsByHouseID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/house/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Невалидные данные ввода", http.StatusBadRequest)
		return
	}

	flats, err := h.FlatService.GetByHouseID(id)
	if err != nil {
		log.Printf("Error getting flats: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	claims, _ := r.Context().Value(models.UserKey).(*models.Claim)
	filteredFlats, err := h.FlatService.FilterByRole(flats, claims.Role)
	if err != nil {
		log.Printf("Error filtering flats: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredFlats)
}
