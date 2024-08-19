package handlers

import (
	"avito-api/internal/avito-api/models"
	"avito-api/internal/avito-api/services"
	"encoding/json"
	"log"
	"net/http"
)

type AuthenticationHandler struct {
	Service *services.AuthenticationService
}

func NewAuthenticationHandler(service *services.AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{Service: service}
}

func (h *AuthenticationHandler) GetDummyJWT(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Метод недоступен", http.StatusMethodNotAllowed)
		return
	}

	userType := r.URL.Query().Get("user_type")

	switch userType {
	case string(models.Moderator), string(models.Client):
		token, err := h.Service.GetDummyJWT(models.UserTypesEnum(userType))
		if err != nil {
			log.Printf("Error getting jwt for dummyLogin: %v", err)
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(token)
	default:
		http.Error(w, "Невалидные данные ввода", http.StatusBadRequest)
		return
	}

}
