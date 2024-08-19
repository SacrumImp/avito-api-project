package middleware

import (
	"avito-api/internal/avito-api/models"
	"avito-api/internal/avito-api/services"
	"context"
	"net/http"
	"strings"
)

func Authenticate(authService *services.AuthenticationService, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Неавторизованный доступ", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := authService.DecodeJWT(tokenStr)
		if err != nil {
			http.Error(w, "Неавторизованный доступ", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), models.UserKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
