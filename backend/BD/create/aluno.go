// BD/create/tables-aluno.go
package create

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func CreateAluno(Ctx context.Context, Client *firestore.Client, novoAluno models.Aluno) error {

	// MELHOR SÓ A MATRICULA DO QUE ESSE ID DINAMICO
	docID := fmt.Sprintf("%s", novoAluno.Matricula)

	_, err := Client.Collection("alunos").Doc(docID).Set(Ctx, novoAluno)
	if err != nil {
		log.Printf("Erro ao criar aluno %s no banco: %v\n", novoAluno.Matricula, err)
		return err
	}

	log.Printf("Aluno %s criado com sucesso! Documento: %s\n", novoAluno.Nome, docID)
	return nil
}
