package update

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	//"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

// UpdateTurma recebe a struct pronta e a envia diretamente para o Firestore
func UpdateVagasOcupadas(Ctx context.Context, Client *firestore.Client, codigoTurma string) error {
	_, err := Client.Collection("turmas_UnB").Doc(codigoTurma).
		Update(Ctx, []firestore.Update{
			{
				Path:  "VagasOcupadas",
				Value: firestore.Increment(1),
			},
		})

	if err != nil {
		log.Printf("Erro ao atualizar vagas: %v\n", err)
		return err
	}

	return nil
}
