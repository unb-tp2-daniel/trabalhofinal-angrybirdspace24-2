// BD/create/turmas.go
package create

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func CreateMateria(Ctx context.Context, Client *firestore.Client, novaMateria models.Materia) error {

	docID := novaMateria.CodigoMateria

	_, err := Client.Collection("materias").Doc(docID).Set(Ctx, novaMateria)

	if err != nil {
		log.Printf("Erro ao criar a materia %s no banco: %v\n", docID, err)
		return err
	}

	log.Printf("Materia %s criada com sucesso no Firestore!\n", docID)
	return nil
}
