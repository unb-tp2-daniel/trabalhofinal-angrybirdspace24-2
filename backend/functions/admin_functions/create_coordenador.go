package admin_functions

import (
	"context"
	//"log"

	"cloud.google.com/go/firestore"

	// Importamos o pacote de modelos
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func CreateCoordenador(Ctx context.Context, Client *firestore.Client, novoCoordenador models.Coordenador) error {
	return nil
}
