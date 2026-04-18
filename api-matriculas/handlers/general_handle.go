package handlers

import (
	"api-matriculas/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"api-matriculas/auth"
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if (r.Method == http.MethodPost) {
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)

		if (err != nil) {
			http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
        	return
		}

		/*
		// checar se o usuário está na database

		if (!checkUser(matricula, senha)) {
			http.Error(w, "Usuário não existe no sistema", http.StatusForbidden)
			return
		}

		*/
		
		token, err := auth.GerarToken(user.Matricula, user.Role)
		if (err != nil) {
			http.Error(w, "Erro ao gerar acesso", http.StatusInternalServerError)
			println(err.Error())
        	return
		}

		// retornar o token pra página
		w.Header().Set("Content-Type", "application/json")
    	json.NewEncoder(w).Encode(map[string]string{"token": token})

		defer r.Body.Close() // defer faz essa função ser chamada apenas dps da função LoginHandle (pai) terminar
		fmt.Printf("novo usuário matricula: %s | senha: %s | role: %s", user.Matricula, user.Senha, user.Role)
		return
	}

	http.Error(w, "Método não permitido", http.StatusForbidden)
}
