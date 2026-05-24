package read

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func GetMatriculaByIds(ctx context.Context, client *firestore.Client, alunoId string, turmaId string) (*models.Matricula, error) {
	var matriculaId string = alunoId + "_" + turmaId;
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

func GetAllMatriculasFromTurma(ctx context.Context, client *firestore.Client, turmaId string) ([]models.Matricula, error){
	var matriculas []models.Matricula

	query := client.Collection("matriculas").Where("TurmaId", "==", turmaId)

	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	for _, doc := range docs {
		var matricula models.Matricula
		err = doc.DataTo(&matricula)

		if err == nil {
			matriculas = append(matriculas, matricula)
		}
	}

	return matriculas, nil
}

func GetAllOrderedMatriculasFromTurma(ctx context.Context, client *firestore.Client, turmaId string) ([]models.Matricula, error){
	var matriculas []models.Matricula

	query := client.Collection("matriculas").
		Where("TurmaId", "==", turmaId).
		OrderBy("PrioridadeNota", firestore.Desc). // o maior prioritario no topo
		OrderBy("dataSolicitacao", firestore.Asc) // desempate por tempo, se ambos tiverem a mesma nota


	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	for _, doc := range docs {
		var matricula models.Matricula
		err = doc.DataTo(&matricula)

		if err == nil {
			matriculas = append(matriculas, matricula)
		}
	}

	return matriculas, nil
}