// BD/create/turmas.go
package create

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func CreateCurso(Ctx context.Context, Client *firestore.Client, novoCurso models.Curso) error {

	docID := novoCurso.CursoId

	_, err := Client.Collection("cursos").Doc(docID).Set(Ctx, novoCurso)

	if err != nil {
		log.Printf("Erro ao criar a curso %s no banco: %v\n", docID, err)
		return err
	}

	log.Printf("Curso %s criado com sucesso no Firestore!\n", docID)
	return nil
}
