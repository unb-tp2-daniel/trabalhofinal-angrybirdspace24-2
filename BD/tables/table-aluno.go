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

func CreateAluno(hash []byte, instituicao string, matricula string,nome string, disciplinas []string) map[string]interface{} {
	alunoTeste := map[string]interface{}{
		"instituicao_id":        instituicao, // Tem que bater com o seu FindInstitutionByID
		"matricula":             matricula,
		"nome":                  nome,
		"senha_hash":            string(hash), // Nunca salve a senha plana!
		"disciplinas_aprovadas": disciplinas,
		"created":               time.Now(),
	}
	return alunoTeste
}
