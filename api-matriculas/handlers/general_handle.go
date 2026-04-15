package handlers

import (
	"api-matriculas/models"
	"encoding/json"
	"net/http"
)

func GenerateStudentToken(w http.ResponseWriter, r *http.Request) {

}

func GenerateTeacherToken(w http.ResponseWriter, r *http.Request) {

	var req models.TeacherToken
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
}
