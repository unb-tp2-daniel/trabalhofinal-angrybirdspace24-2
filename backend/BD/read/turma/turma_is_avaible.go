package turma

import (
	"context"

	"cloud.google.com/go/firestore"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func TurmaIsAvailable(Ctx context.Context, Client *firestore.Client, id string) (bool, error) {

	doc, err := Client.Collection("turmas_UnB").Doc(id).Get(Ctx)

	if err != nil {
		return false, err
	}

	var turma models.Turma

	err = doc.DataTo(&turma)

	if err != nil {
		return false, err
	}

	disponivel := turma.VagasOcupadas < turma.VagasTotais

	return disponivel, nil
}
