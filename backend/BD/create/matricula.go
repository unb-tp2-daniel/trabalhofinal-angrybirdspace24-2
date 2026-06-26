// BD/create/tables-aluno.go
package create

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateMatricula(Ctx context.Context, Client *firestore.Client, novaMatricula models.Matricula) error {
	docID := fmt.Sprintf("%s_%s", novaMatricula.AlunoId, novaMatricula.TurmaId)

	_, err := Client.Collection("matriculas").Doc(docID).Create(Ctx, novaMatricula)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			log.Printf("Aluno %s já possui solicitação para a turma %s\n", novaMatricula.AlunoId, novaMatricula.TurmaId)
			return fmt.Errorf("aluno_ja_matriculado")
		}
		
		log.Printf("Erro ao criar matrícula para aluno %s no banco: %v\n", novaMatricula.AlunoId, err)
		return err
	}

	log.Printf("Matrícula para aluno %s criada com sucesso! Documento: %s\n", novaMatricula.AlunoId, docID)
	return nil
}
