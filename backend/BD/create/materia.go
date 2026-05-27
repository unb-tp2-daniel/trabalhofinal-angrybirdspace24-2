package create

import (
	"context"
	//"log"

	"cloud.google.com/go/firestore"

	// Importamos o pacote de modelos
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func CreateMateria(Ctx context.Context, Client *firestore.Client, novaMateria models.Materia) error {
	return nil
}
