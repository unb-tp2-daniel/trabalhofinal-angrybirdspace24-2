// BD/create/turmas.go
package create

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func CreateProfessor(Ctx context.Context, Client *firestore.Client, novoProfessor models.Professor) error {

	docID := novoProfessor.ProfessorId

	_, err := Client.Collection("professores").Doc(docID).Set(Ctx, novoProfessor)

	if err != nil {
		log.Printf("Erro ao criar a professor %s no banco: %v\n", docID, err)
		return err
	}

	log.Printf("Professor %s criado com sucesso no Firestore!\n", docID)
	return nil
}
