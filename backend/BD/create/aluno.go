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
func CreateAluno(Ctx context.Context, Client *firestore.Client, novoAluno models.Aluno) error {
	// Gera um ID dinâmico ("Unb_20260001", "Unb_20260002")
	docID := fmt.Sprintf("Unb_%s", novoAluno.Matricula)
	// O Firestore lê as tags da struct e faz o mapeamento sozinho
	_, err := Client.Collection("alunos").Doc(docID).Set(Ctx, novoAluno)
	if err != nil {
		log.Printf("Erro ao criar aluno %s no banco: %v\n", novoAluno.Matricula, err)
		return err
	}

	log.Printf("Aluno %s criado com sucesso! Documento: %s\n", novoAluno.Nome, docID)
	return nil
}
