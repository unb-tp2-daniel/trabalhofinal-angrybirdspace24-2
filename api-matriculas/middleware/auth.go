package middleware

import (
	"net/http"
	"os"
)

func RequireInstitutionKey(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("School-API-Key")
		expectedKey := os.Getenv("INSTITUTION_KEY")
		if expectedKey == "" {
			expectedKey = "coco-bosta-peluda" // valor da chave para desenvolvimento local, deve ser sobrescrito em produção
		}
		if key != expectedKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
