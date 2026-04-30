package database

import (
	"context"

	//"time"

	"log"

	"cloud.google.com/go/firestore"
)

func CreateTurma(Ctx context.Context, Client *firestore.Client, hash []byte, codigoTurma string, materiaTurma string, nomeMateria string, semestre string, capacidadeTurma int, vagasOcupadas int, vagasRestantes int, ativa bool) map[string]interface{} {
	turmaTeste := map[string]interface{}{
		"codigoTurma":    codigoTurma,
		"materiaId":      materiaTurma,
		"nomeMateria":    nomeMateria,
		"semestre":       semestre,
		"capacidade":     capacidadeTurma,
		"ocupadas":       vagasOcupadas,
		"vagasRestantes": vagasRestantes,
		"status":         ativa,
	}
	// 3. Salvando no Firestore com um ID Composto (Instituicao + Matricula)
	docID := "Unb_20260001"
	_, err := Client.Collection("turmas_UnB").Doc(codigoTurma).Set(Ctx, turmaTeste)

	if err != nil {
		log.Println("Erro ao criar turma", err)
	} else {
		log.Printf("Turma crida com sucesso: %s\n", docID)
	}
	return turmaTeste
}
