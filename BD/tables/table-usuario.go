package database

import (
	"context"

	"time"

	"log"

	"cloud.google.com/go/firestore"
)

// Client vai guardar a conexão com o banco para a API usar

func CreateUsuario(Ctx context.Context, 
	Client *firestore.Client,
	matricula string,
	senha string,
	nivelAcesso string,
	) map[string]interface{} {

	// criar um curso
	var err error
	usuario := map[string]interface{}{
		"matricula":    matricula,
		"senha":   senha,
		"departamentoid":    nivelAcesso,
		"created": time.Now(),
	}
	_, err = Client.Collection("usuarios").Doc(matricula).Set(Ctx, usuario)
	if err != nil {
		log.Println("erro ao criar usuario:", err)
	}else{
		log.Println("usuario criado com sucesso!")
	}
	return usuario

}
