// BD/create/turmas.go
package create

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"

	// Importamos o pacote de modelos
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

// CreateTurma recebe a struct pronta e a envia diretamente para o Firestore
func CreateProfessor(Ctx context.Context, Client *firestore.Client, novoProfessor models.Professor) error {

	// Usamos o próprio código da turma como ID do documento
	docID := novoProfessor.ProfessorId

	// O Firestore lê as tags da struct e salva no banco
	_, err := Client.Collection("professores").Doc(docID).Set(Ctx, novoProfessor)

	// Verificamos se houve erro e avisamos no terminal
	if err != nil {
		log.Printf("Erro ao criar a professor %s no banco: %v\n", docID, err)
		return err
	}

	log.Printf("Professor %s criado com sucesso no Firestore!\n", docID)
	return nil
}
