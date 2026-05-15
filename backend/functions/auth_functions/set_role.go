package auth_functions

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go/v4"
)

func SetRole(w http.ResponseWriter, r *http.Request) {

	app, err := firebase.NewApp(context.Background(), nil)

	if err != nil {
		http.Error(w, "erro firebase", 500)
		return
	}

	authClient, err := app.Auth(context.Background())

	if err != nil {
		http.Error(w, "erro auth", 500)
		return
	}

	uid := r.URL.Query().Get("uid")
	role := r.URL.Query().Get("role")

	if uid == "" || role == "" {
		http.Error(w, "uid e role obrigatórios", 400)
		return
	}

	err = authClient.SetCustomUserClaims(
		context.Background(),
		uid,
		map[string]interface{}{
			"role": role,
		},
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("role definida"))
}
