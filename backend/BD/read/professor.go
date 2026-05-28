package read

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

// GetAllTurmas faz a busca diretamente no Firestore e retorna uma lista de turmas.
// Retorna a lista pronta ou um erro, se houver.
func GetAllProfessores(Ctx context.Context, Client *firestore.Client) ([]models.Professor, error) {
	var professores []models.Professor

	// Faz a busca na coleção do Firestore
	iter := Client.Collection("professores").Documents(Ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break // Fim da leitura
		}
		if err != nil {
			return nil, err // Retorna o erro para o Handler tratar
		}

		var p models.Professor
		if err := doc.DataTo(&p); err != nil {
			// Se der erro ao converter um documento específico, apenas ignora e vai pro próximo
			continue
		}

		professores = append(professores, p)
	}

	return professores, nil
}
