package database

import (
	"context"

	"time"

	"log"

	"cloud.google.com/go/firestore"
)

// Client vai guardar a conexão com o banco para a API usar

func CreateDepartamento(Ctx context.Context, 
	Client *firestore.Client,
	coordenadorId string,
	departamentoNome string,
	departamentoiId string,
	) map[string]interface{} {

	// criar um curso
	var err error
	departamento := map[string]interface{}{
		"departamentoId":		departamentoiId,
		"departamentoNome":		departamentoNome,
		"corrdenadorId":		coordenadorId,
		"created": time.Now(),
	}
	_, err = Client.Collection("professores").Doc(departamentoiId).Set(Ctx, departamento)
	if err != nil {
		log.Println("erro ao criar departamento:", err)
	}else{
		log.Println("drpartamento criado com sucesso!")
	}
	return departamento

}
