package handlers

import (
	"api-matriculas/models"
	"encoding/json"
	"net/http"
	"time"
)

func GenerateStudentToken(w http.ResponseWriter, r *http.Request) {
	var req models.StudentAuth
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	tokenData := models.StudentToken{
		StudentID:     req.StudentID,
		InstitutionID: req.InstitutionID,
		ExpiresAt:     time.Now().Add(30 * time.Minute),
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tokenData)
	if err != nil {
		http.Error(w, "Erro ao retornar token", http.StatusInternalServerError)
		return
	}
}

func GenerateTeacherToken(w http.ResponseWriter, r *http.Request) {

	var req models.TeacherAuth
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	tokenData := models.TeacherToken{
		TeacherID:     req.TeacherID,
		InstitutionID: req.InstitutionID,
		ExpiresAt:     time.Now().Add(30 * time.Minute),
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tokenData)
	if err != nil {
		http.Error(w, "Erro ao retornar token", http.StatusInternalServerError)
		return
	}
}
