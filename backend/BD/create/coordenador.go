// BD/create/turmas.go
package create

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"

	// Importamos o pacote de modelos
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func CreateCoordenador(Ctx context.Context, Client *firestore.Client, novoCoordenador models.Coordenador) error {

	docID := novoCoordenador.CoordenadorId

	_, err := Client.Collection("coordenadores").Doc(docID).Set(Ctx, novoCoordenador)

	if err != nil {
		log.Printf("Erro ao criar a coordenador %s no banco: %v\n", docID, err)
		return err
	}

	log.Printf("Coordenador %s criado com sucesso no Firestore!\n", docID)
	return nil
}
