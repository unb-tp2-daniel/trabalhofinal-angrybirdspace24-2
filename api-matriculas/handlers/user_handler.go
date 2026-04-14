package handlers

import (
	"api-matriculas/models"
	"encoding/json"
	"net/http"
)

var Users []models.User

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		for _, u := range Users {
			w.Write([]byte(u.Nome + "\n"))
		}
		return
	}
	if r.Method == http.MethodPost {
		var u models.User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, "json invalido", http.StatusBadRequest)
			return
		}
		Users = append(Users, u)
		w.Write([]byte("POST recebido"))
		return
	}
	http.Error(w, "metodo nao permitido", http.StatusMethodNotAllowed)
}
