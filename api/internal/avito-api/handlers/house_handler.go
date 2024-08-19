package handlers

import (
	"avito-api/internal/avito-api/models"
	"avito-api/internal/avito-api/services"
	"encoding/json"
	"log"
	"net/http"
)

type HouseHandler struct {
	Service *services.HouseService
}

func NewHouseHandler(service *services.HouseService) *HouseHandler {
	return &HouseHandler{Service: service}
}

func (h *HouseHandler) CreateHouse(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод недоступен", http.StatusMethodNotAllowed)
		return
	}

	var houseInputObject models.HouseInputObject
	if err := json.NewDecoder(r.Body).Decode(&houseInputObject); err != nil {
		http.Error(w, "Невалидные данные ввода", http.StatusBadRequest)
		return
	}

	if houseInputObject.Address == "" || houseInputObject.Year <= 0 {
		http.Error(w, "Невалидные данные ввода", http.StatusBadRequest)
		return
	}

	house, err := h.Service.CreateHouse(&houseInputObject)
	if err != nil {
		log.Printf("Error inserting house: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(house)
}
