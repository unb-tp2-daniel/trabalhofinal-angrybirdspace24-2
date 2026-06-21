package matricula

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func GetAllMatriculasFromTurma(ctx context.Context, client *firestore.Client, turmaId string) ([]models.Matricula, error) {
	var matriculas []models.Matricula

	query := client.Collection("matriculas").Where("TurmaId", "==", turmaId)

	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	for _, doc := range docs {
		var matricula models.Matricula
		err = doc.DataTo(&matricula)

		if err == nil {
			matriculas = append(matriculas, matricula)
		}
	}

	return matriculas, nil
}
