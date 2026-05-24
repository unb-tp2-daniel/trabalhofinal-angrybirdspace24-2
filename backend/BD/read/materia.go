package read

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func GetMateriaById(ctx context.Context, client *firestore.Client, id string) (*models.Materia, error) {
	docSnap, err := client.Collection("materias").Doc(id).Get(ctx)

	if err != nil {
		return nil, fmt.Errorf("matéria %s não encontrado na busca: %v", id, err)
	}

	var materia models.Materia
	err = docSnap.DataTo(&materia)

	if err != nil {
		return nil, fmt.Errorf("erro ao converter dados do curso: %v", err)
	}

	return &materia, nil
}