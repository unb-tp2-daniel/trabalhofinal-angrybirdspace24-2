package coordenador

import (
	"context"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func GetAllCoordenadores(Ctx context.Context, Client *firestore.Client) ([]models.Coordenador, error) {
	var alunos []models.Coordenador

	iter := Client.Collection("coordenadores").Documents(Ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var t models.Coordenador
		if err := doc.DataTo(&t); err != nil {

			continue
		}

		alunos = append(alunos, t)
	}

	return alunos, nil
}