package read

import (
	"context"
	"fmt"
	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

/*
Essa função serve para caso o projeto siga o padrão atual: Instituição_matricula
Caso mude para o padrão de todas as outras coleções, o código comentado abaixo desse é o ideal
*/

func GetAlunoById(ctx context.Context, client *firestore.Client, matricula string) (*models.Aluno, error) {
	query := client.Collection("alunos").Where("matricula", "==", matricula).Limit(1)

	doc, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	// curso não existe
	if len(doc) == 0 {
		return nil, fmt.Errorf("aluno de matricula %s não encontrado", matricula)
	}

	var aluno models.Aluno
	err = doc[0].DataTo(&aluno)

	if err != nil {
		return nil, err
	}

	return &aluno, nil
}

func GetAllAlunos(Ctx context.Context, Client *firestore.Client) ([]models.Aluno, error) {
	var alunos []models.Aluno

	// Faz a busca na coleção do Firestore
	iter := Client.Collection("alunos").Documents(Ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break // Fim da leitura
		}
		if err != nil {
			return nil, err // Retorna o erro para o Handler tratar
		}

		var t models.Aluno
		if err := doc.DataTo(&t); err != nil {
			// Se der erro ao converter um documento específico, apenas ignora e vai pro próximo
			continue
		}

		alunos = append(alunos, t)
	}

	return alunos, nil
}

/*
func GetAlunoById(ctx context.Context, client *firestore.Client, id string) (*models.Aluno, error) {
	docSnap, err := client.Collection("alunos").Doc(id).Get(ctx)

	if err != nil {
		return nil, fmt.Errorf("aluno %s não encontrado na busca: %v", id, err)
	}

	var aluno models.Aluno
	err = docSnap.DataTo(&aluno)

	if err != nil {
		return nil, fmt.Errorf("erro ao converter dados do curso: %v", err)
	}

	return &aluno, nil
}*/