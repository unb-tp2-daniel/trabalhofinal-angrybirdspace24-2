// BD/create/turmas.go
package create

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func CreateDepartamento(Ctx context.Context, Client *firestore.Client, novoDepartamento models.Departamento) error {

	docID := novoDepartamento.DepartamentoId

	_, err := Client.Collection("departamentos").Doc(docID).Set(Ctx, novoDepartamento)

	if err != nil {
		log.Printf("Erro ao criar a departamento %s no banco: %v\n", docID, err)
		return err
	}

	log.Printf("Departamento %s criado com sucesso no Firestore!\n", docID)
	return nil
}
