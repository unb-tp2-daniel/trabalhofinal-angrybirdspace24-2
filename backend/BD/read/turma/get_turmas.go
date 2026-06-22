package turma

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

// GetAllTurmas faz a busca diretamente no Firestore e retorna uma lista de turmas.
// Retorna a lista pronta ou um erro, se houver.
func GetTurmas(Ctx context.Context, Client *firestore.Client, max_results int) ([]models.Turma, error) {
	var turmas []models.Turma

	// Faz a busca na coleção do Firestore
	iter := Client.Collection("turmas").Limit(max_results).Documents(Ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break // Fim da leitura
		}
		if err != nil {
			return nil, err // Retorna o erro para o Handler tratar
		}

		var t models.Turma
		if err := doc.DataTo(&t); err != nil {
			// Se der erro ao converter um documento específico, apenas ignora e vai pro próximo
			continue
		}

		turmas = append(turmas, t)
	}

	return turmas, nil
}
