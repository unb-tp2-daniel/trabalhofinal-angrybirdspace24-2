package create

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/models"
)

func CreateTurma(Ctx context.Context, Client *firestore.Client, novaTurma models.Turma) error {
	_, err := Client.Collection("turmas_UnB").Doc(novaTurma.CodigoTurma).Set(Ctx, novaTurma)
	return err
}
