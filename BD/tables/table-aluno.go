package database

import (
	"context"
	
	"time"

	"cloud.google.com/go/firestore"
)

// Client vai guardar a conexão com o banco para a API usar
var (
	Client *firestore.Client
	Ctx    context.Context
)

func CreateAluno(hash []byte) map[string]interface{} {
	alunoTeste := map[string]interface{}{
		"instituicao_id":        "Unb", // Tem que bater com o seu FindInstitutionByID
		"matricula":             "20260001",
		"nome":                  "Guilherme Silva Cavalcantessss",
		"senha_hash":            string(hash), // Nunca salve a senha plana!
		"disciplinas_aprovadas": []string{"MAT0011"},
		"created":               time.Now(),
	}
	return alunoTeste
}
