package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"trabalho/api-matriculas/models"
	"trabalho/api-matriculas/repository"
	"trabalho/api-matriculas/services"
)

func GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	var req models.Auth
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	tokenData, err := services.GenerateAccessToken(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tokenData)
	if err != nil {
		http.Error(w, "Erro ao retornar token", http.StatusInternalServerError)
		return
	}

}

func VerifyKeyHandler(w http.ResponseWriter, r *http.Request) {
	institutionalKey := r.Header.Get("institutional_key")
	institution, err := repository.FindInstitutionByID(institutionalKey)
	if err != nil {
		http.Error(w, "Chave institucional inválida", http.StatusUnauthorized)
		return
	}
	w.Write([]byte("Chave institucional válida para a instituição: " + institution))
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Online. " + r.Method + " recebido."))
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC1123)
	w.Write([]byte("Current time: " + currentTime))
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Status: Online"))
}
