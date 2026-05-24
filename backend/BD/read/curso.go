package read

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func GetCursoById(ctx context.Context, client *firestore.Client, id string) (*models.Curso, error) {
	docSnap, err := client.Collection("cursos").Doc(id).Get(ctx)

	if err != nil {
		return nil, fmt.Errorf("curso %s não encontrado na busca: %v", id, err)
	}

	var curso models.Curso
	err = docSnap.DataTo(&curso)

	if err != nil {
		return nil, fmt.Errorf("erro ao converter dados do curso: %v", err)
	}

	return &curso, nil
}