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
func CreateCoordenador(Ctx context.Context, Client *firestore.Client, novoCoordenador models.Coordenador) error {

	// Usamos o próprio código da turma como ID do documento
	docID := novoCoordenador.CoordenadorId

	// O Firestore lê as tags da struct e salva no banco
	_, err := Client.Collection("coordenadores").Doc(docID).Set(Ctx, novoCoordenador)

	// Verificamos se houve erro e avisamos no terminal
	if err != nil {
		log.Printf("Erro ao criar a coordenador %s no banco: %v\n", docID, err)
		return err
	}

	log.Printf("Coordenador %s criado com sucesso no Firestore!\n", docID)
	return nil
}
