package auth_functions

import (
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"net/http"
	"context"
)

func CreateUser(w http.ResponseWriter, r *http.Request){
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

	matricula := r.URL.Query().Get("matricula")
	email := r.URL.Query().Get("email")
	senha := r.URL.Query().Get("senha")

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