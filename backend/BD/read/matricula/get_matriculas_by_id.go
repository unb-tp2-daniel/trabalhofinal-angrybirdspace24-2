package matricula

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func GetMatriculaByIds(ctx context.Context, client *firestore.Client, alunoId string, turmaId string) (*models.Matricula, error) {
	var matriculaId string = alunoId + "_" + turmaId
	docSnap, err := client.Collection("matriculas").Doc(matriculaId).Get(ctx)

	if err != nil {
		return nil, fmt.Errorf("matricula %s não encontrado na busca: %v", matriculaId, err)
	}

	var matricula models.Matricula
	err = docSnap.DataTo(&matricula)

	if err != nil {
		return nil, fmt.Errorf("erro ao converter dados do curso: %v", err)
	}

	return &matricula, nil
}
