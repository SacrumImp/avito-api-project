package handlers

import (
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
		http.Error(w, "Failed to retrieve houses", http.StatusBadRequest)
		return
	}
	flats, err := h.Service.GetByHouseID(id)
	if err != nil {
		log.Printf("Error getting flats: %v", err)
		http.Error(w, "Failed to retrieve flats", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flats)
}
