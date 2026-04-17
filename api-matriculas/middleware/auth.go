package middleware

import (
	"net/http"
	"os"
)

func RequireInstitutionKey(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("institutional_key")
		expectedKey := os.Getenv("INSTITUTION_KEY")
		if expectedKey == "" {
			http.Error(w, "Service unavailable", http.StatusInternalServerError)
		}
		if key != expectedKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
