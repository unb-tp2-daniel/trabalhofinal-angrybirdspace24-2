package auth_functions

import (
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"net/http"
	"context"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
	"log"
	"encoding/json"
)

func CreateUser(w http.ResponseWriter, r *http.Request){
	//Tornando o acesso visível para o front
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}

	context := context.Background()

	app, err := firebase.NewApp(context, nil)
	if err != nil {
		http.Error(w, "Erro ao chamar firebase", 500)
		return
	}

	authClient, err := app.Auth(context)
	if err != nil {
		http.Error(w, "Erro com o auth", 500)
		return
	}

	var usuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		log.Printf("Erro ao decodificar JSON do usuário: %v", err)
		http.Error(w, "Formato de dados inválido", http.StatusBadRequest)
		return
	}

	matricula := usuario.Matricula
	email := usuario.Email
	senha := usuario.Senha

	if senha == "" || email == "" || matricula == ""{
		http.Error(w, "Campos incompletos", 400)
		return
	}

	params := (&auth.UserToCreate{}).
		UID(matricula).
		Email(email).
		Password(senha)

	userRecord, err := authClient.CreateUser(context, params)

	if err != nil {
		http.Error(w, "Erro ao criar usuário", 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(userRecord.UID))
}