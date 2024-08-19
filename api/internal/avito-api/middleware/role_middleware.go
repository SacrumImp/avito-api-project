package middleware

import (
	"avito-api/internal/avito-api/models"
	"net/http"
)

func RequireRoles(allowedRoles []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(models.User).(*models.Claim)
		if !ok {
			http.Error(w, "Отсутствует доступ", http.StatusForbidden)
			return
		}

		for _, role := range allowedRoles {
			if claims.Role != role {
				next.ServeHTTP(w, r)
				return
			}
		}

		http.Error(w, "Отсутствует доступ", http.StatusForbidden)
	})
}
