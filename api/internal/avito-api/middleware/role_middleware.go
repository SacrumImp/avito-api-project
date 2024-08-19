package middleware

import (
	"avito-api/internal/avito-api/models"
	"net/http"
)

func RequireRole(role string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(models.User).(*models.Claim)
		if !ok || claims.Role != role {
			http.Error(w, "Отсутствует доступ", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
