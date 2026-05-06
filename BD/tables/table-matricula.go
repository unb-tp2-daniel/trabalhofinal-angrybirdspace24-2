package database

import (
	"context"

	"time"

	"log"

	"cloud.google.com/go/firestore"
)

func CreateMatricula(Ctx context.Context, Client *firestore.Client,
	alunoId string,
	turmaId string,
	status bool,
	dataSolicitacao time.Time,
	prioridades map[string]string,
	semestre string,
) map[string]interface{} {
	matriculaTeste := map[string]interface{}{
		"alunoId":         alunoId,
		"turmaId":         turmaId,
		"status":          status,
		"dataSolicitacao": dataSolicitacao,
		"prioridades":     prioridades,
		"semestre":        semestre,
	}
	_, err := Client.Collection("matriculas_UnB").Doc(alunoId).Set(Ctx, matriculaTeste)

	if err != nil {
		log.Println("Erro ao criar turma", err)
	} else {
		log.Printf("Turma crida com sucesso")
	}
	return matriculaTeste
}
