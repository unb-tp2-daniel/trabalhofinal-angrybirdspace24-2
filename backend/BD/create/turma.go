// BD/create/turmas.go
package create

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func CreateTurma(Ctx context.Context, Client *firestore.Client, novaTurma models.Turma) error {

	docID := novaTurma.CodigoTurma

	_, err := Client.Collection("turmas").Doc(docID).Set(Ctx, novaTurma)

	if err != nil {
		log.Printf("Erro ao criar a turma %s no banco: %v\n", docID, err)
		return err
	}

	log.Printf("Turma %s criada com sucesso no Firestore!\n", docID)
	return nil
}
