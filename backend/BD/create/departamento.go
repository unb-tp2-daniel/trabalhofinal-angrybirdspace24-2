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
func CreateDepartamento(Ctx context.Context, Client *firestore.Client, novoDepartamento models.Departamento) error {

	// Usamos o próprio código da turma como ID do documento
	docID := novoDepartamento.DepartamentoId

	// O Firestore lê as tags da struct e salva no banco
	_, err := Client.Collection("departamentos").Doc(docID).Set(Ctx, novoDepartamento)

	// Verificamos se houve erro e avisamos no terminal
	if err != nil {
		log.Printf("Erro ao criar a departamento %s no banco: %v\n", docID, err)
		return err
	}

	log.Printf("Departamento %s criado com sucesso no Firestore!\n", docID)
	return nil
}
