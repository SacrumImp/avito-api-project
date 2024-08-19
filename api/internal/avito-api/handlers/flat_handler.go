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

type FlatHandler struct {
	Service *services.FlatService
}

func NewFlatHandler(service *services.FlatService) *FlatHandler {
	return &FlatHandler{Service: service}
}

func (h *FlatHandler) GetByHouseID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/house/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Невалидные данные ввода", http.StatusBadRequest)
		return
	}
	flats, err := h.Service.GetByHouseID(id)
	if err != nil {
		log.Printf("Error getting flats: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flats)
}

func (h *FlatHandler) CreateFlat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод недоступен", http.StatusMethodNotAllowed)
		return
	}

	var flatInputObject models.FlatInputObject
	if err := json.NewDecoder(r.Body).Decode(&flatInputObject); err != nil {
		http.Error(w, "Невалидные данные ввода", http.StatusBadRequest)
		return
	}

	if flatInputObject.HouseId <= 0 || flatInputObject.Price < 0 || flatInputObject.Rooms <= 0 {
		http.Error(w, "Невалидные данные ввода", http.StatusBadRequest)
		return
	}

	flat, err := h.Service.CreateFlat(&flatInputObject)
	if err != nil {
		log.Printf("Error inserting flat: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flat)
}

func (h *FlatHandler) UpdateFlatStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод недоступен", http.StatusMethodNotAllowed)
		return
	}

	var flatUpdateObject models.FlatUpdateObject
	if err := json.NewDecoder(r.Body).Decode(&flatUpdateObject); err != nil || flatUpdateObject.HouseId <= 0 || flatUpdateObject.FlatId <= 0 || flatUpdateObject.Status == "" {
		http.Error(w, "Невалидные данные ввода", http.StatusBadRequest)
		return
	}

	flat, err := h.Service.UpdateFlatStatus(&flatUpdateObject)
	if err != nil {
		log.Printf("Error updating flat: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flat)

}
