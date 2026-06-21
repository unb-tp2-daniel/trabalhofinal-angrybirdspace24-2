package turma

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func GetTurmaById(ctx context.Context, client *firestore.Client, id string) (*models.Turma, error) {
	docSnap, err := client.Collection("turmas").Doc(id).Get(ctx)

	if err != nil {
		return nil, fmt.Errorf("turma %s não encontrado na busca: %v", id, err)
	}

	var turma models.Turma
	err = docSnap.DataTo(&turma)

	if err != nil {
		return nil, fmt.Errorf("erro ao converter dados do curso: %v", err)
	}

	return &turma, nil
}
