package materia

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
	"google.golang.org/api/iterator"
)

func GetAllMaterias(Ctx context.Context, Client *firestore.Client) ([]models.Materia, error) {
	var materias []models.Materia

	// Faz a busca na coleção do Firestore
	iter := Client.Collection("materias").Documents(Ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var m models.Materia
		if err := doc.DataTo(&m); err != nil {

			continue
		}

		materias = append(materias, m)
	}

	return materias, nil
}
