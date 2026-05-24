package read

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

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

// GetAllTurmas faz a busca diretamente no Firestore e retorna uma lista de turmas.
// Retorna a lista pronta ou um erro, se houver.
func GetAllTurmas(Ctx context.Context, Client *firestore.Client) ([]models.Turma, error) {
	var turmas []models.Turma

	// Faz a busca na coleção do Firestore
	iter := Client.Collection("turmas_UnB").Documents(Ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break // Fim da leitura
		}
		if err != nil {
			return nil, err // Retorna o erro para o Handler tratar
		}

		var t models.Turma
		if err := doc.DataTo(&t); err != nil {
			// Se der erro ao converter um documento específico, apenas ignora e vai pro próximo
			continue
		}

		turmas = append(turmas, t)
	}

	return turmas, nil
}
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
