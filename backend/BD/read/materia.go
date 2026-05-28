package read

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
	"google.golang.org/api/iterator"
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

func GetAllMaterias(Ctx context.Context, Client *firestore.Client) ([]models.Materia, error) {
	var materias []models.Materia

	// Faz a busca na coleção do Firestore
	iter := Client.Collection("materias").Documents(Ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break // Fim da leitura
		}
		if err != nil {
			return nil, err // Retorna o erro para o Handler tratar
		}

		var m models.Materia
		if err := doc.DataTo(&m); err != nil {
			// Se der erro ao converter um documento específico, apenas ignora e vai pro próximo
			continue
		}

		materias = append(materias, m)
	}

	return materias, nil
}
