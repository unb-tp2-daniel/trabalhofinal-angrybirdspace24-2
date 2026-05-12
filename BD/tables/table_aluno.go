package database

import (
	"context"

	//"time"

	"log"

	"cloud.google.com/go/firestore"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/BD/enums"
)

// Client vai guardar a conexão com o banco para a API usar

func CreateAluno(Ctx context.Context, Client *firestore.Client, hash []byte,
	matricula string,
	cursoId string,
	ativo bool,
	semestre string,
	materiasConcluidas map[string]string,
	ira float64,
	prioridades map[string]string,
	nivelAcademico enums.NivelAcademico,
	nomeAluno string) map[string]interface{} {
	alunoTeste := map[string]interface{}{
		"matricula":          matricula,
		"cursoId":            cursoId,
		"ativo":              ativo,
		"semestre":           semestre,
		"materiasConcluidas": materiasConcluidas,
		"ira":                ira,
		"prioridades":        prioridades,
		"nivelAcademico":     nivelAcademico,
	}
	// 3. Salvando no Firestore com um ID Composto (Instituicao + Matricula)
	docID := "Unb_20260001"
	_, err := Client.Collection("alunos").Doc(docID).Set(Ctx, alunoTeste)

	if err != nil {
		log.Println("Erro ao criar aluno de teste:", err)
	} else {
		log.Printf("Aluno de teste criado com sucesso! Documento: %s\n", docID)
	}
	return alunoTeste
}
