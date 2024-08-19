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

func (h *AuthenticationHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод недоступен", http.StatusMethodNotAllowed)
		return
	}

	var userRegister models.UserRegisterObject
	if err := json.NewDecoder(r.Body).Decode(&userRegister); err != nil || userRegister.Email == "" || userRegister.Password == "" || userRegister.UserType == "" {
		http.Error(w, "Невалидные данные ввода", http.StatusBadRequest)
		return
	}

	userLogin, err := h.Service.CreateUser(&userRegister)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userLogin)
}

func (h *AuthenticationHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод недоступен", http.StatusMethodNotAllowed)
		return
	}

	var userLogin models.UserLoginObject
	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil || userLogin.UserId <= 0 || userLogin.Password == "" {
		http.Error(w, "Невалидные данные ввода", http.StatusBadRequest)
		return
	}

	token, err := h.Service.LoginUser(&userLogin)
	if err != nil {
		log.Printf("Error authentication user: %v", err)
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
