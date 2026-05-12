package database

import (
	"context"

	"time"

	"log"

	"cloud.google.com/go/firestore"
)

// Client vai guardar a conexão com o banco para a API usar

func CreateProfessor(Ctx context.Context,
	Client *firestore.Client,
	professorid string,
	professorNome string,
	departamentoid string,
) map[string]interface{} {

	// criar um curso confirma
	var err error
	professor := map[string]interface{}{
		"professorid":    professorid,
		"professorNome":  professorNome,
		"departamentoid": departamentoid,
		"created":        time.Now(),
	}
	_, err = Client.Collection("professores").Doc(professorid).Set(Ctx, professor)
	if err != nil {
		log.Println("erro ao criar professor:", err)
	} else {
		log.Println("professor criado com sucesso!")
	}
	return professor

}
