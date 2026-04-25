package middleware

import (
	"net/http"
	"trabalho/api-matriculas/repository"
)

func RequireInstitutionKey(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("institutional_key")
		_, err := repository.FindInstitutionByID(key)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
