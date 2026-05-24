// BD/create/tables-aluno.go
package create

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

// CreateAluno recebe a struct pronta e a envia diretamente para o Firestore
func CreateMatricula(Ctx context.Context, Client *firestore.Client, novaMatricula models.Matricula) error {
	// Gera um ID dinâmico ("Unb_20260001", "Unb_20260002")
	
	/* imagino que um id do tipo alunoId_turmaId seja mais descritivo e melhor posteriormente (por exemplo, para resgatar a posição atual de um único aluno na turma)*/
	//docID := fmt.Sprintf("Unb_%s", novaMatricula.AlunoId)
	docID := fmt.Sprintf("%s_%s", novaMatricula.AlunoId, novaMatricula.TurmaId)

	// O Firestore lê as tags da struct e faz o mapeamento sozinho
	_, err := Client.Collection("matriculas").Doc(docID).Set(Ctx, novaMatricula)
	if err != nil {
		log.Printf("Erro ao criar matrícula para aluno %s no banco: %v\n", novaMatricula.AlunoId, err)
		return err
	}

	log.Printf("Matrícula para aluno %s criada com sucesso! Documento: %s\n", novaMatricula.AlunoId, docID)
	return nil
}
