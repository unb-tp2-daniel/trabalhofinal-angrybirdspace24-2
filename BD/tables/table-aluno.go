package database

import (
	"context"

	"time"

	"log"

	"cloud.google.com/go/firestore"
)

// Client vai guardar a conexão com o banco para a API usar

func CreateAluno(Ctx context.Context, Client *firestore.Client, hash []byte, instituicao string, matricula string, nome string, disciplinas []string) map[string]interface{} {
	alunoTeste := map[string]interface{}{
		// Tem que bater com o seu FindInstitutionByID
		"matricula":             matricula,
		"nome":                  nome,
		"senha_hash":            string(hash), // Nunca salve a senha plana!
		"disciplinas_aprovadas": disciplinas,
		"created":               time.Now(),
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
