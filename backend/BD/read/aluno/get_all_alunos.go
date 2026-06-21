package aluno

import (
	"context"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func GetAllAlunos(Ctx context.Context, Client *firestore.Client) ([]models.Aluno, error) {
	var alunos []models.Aluno

	iter := Client.Collection("alunos").Documents(Ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var t models.Aluno
		if err := doc.DataTo(&t); err != nil {

			continue
		}

		alunos = append(alunos, t)
	}

	return alunos, nil
}
