package auth_functions

import (
	"context"
	"encoding/json"
	"net/http"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/iterator"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"

)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido. Use GET.", http.StatusMethodNotAllowed)
		return
	}

	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("erro no firebase: %v", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("erro no cliente: %v", err)
	}

	var usuarios []models.Usuario

	iter := client.Users(ctx, "") 
	for {
		user, err := iter.Next() 
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Erro ao recuperar usuário: %v", err)
		}

		usuarios = append(usuarios, models.Usuario{
			Matricula: user.UID,
			Email: user.Email,
			Senha: "",
		})
	}
	
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		log.Printf("Erro ao retornar JSON: %v", err)
		http.Error(w, "Erro na resposta", http.StatusInternalServerError)
	}
}