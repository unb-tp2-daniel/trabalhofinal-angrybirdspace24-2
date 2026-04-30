package database

import (
	"context"
	"log"
	"time"

	table_aluno "trabalho/BD/tables"

	"cloud.google.com/go/firestore"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/option"
)

// Client vai guardar a conexão com o banco para a API usar
var (
	Client *firestore.Client
	Ctx    context.Context
)

func InitDB() {
	Ctx = context.Background()

	opt := option.WithAuthCredentialsFile(option.ServiceAccount, "../serviceAccountKey.json")

	//Conecta diretamente no banco matriculas242
	var err error
	Client, err = firestore.NewClientWithDatabase(Ctx, "matriculas242", "matriculas242", opt)

	if err != nil {
		log.Fatalf("Erro ao conectar no Firestore: %v", err)
	}

	log.Println(" BD online")
}

// SeedBaseData cria dados iniciais no banco
func SeedBaseData() {
	log.Println("Iniciando seed do banco...")

	// 1. Criptografando a senha de teste ANTES de salvar no banco
	senhaPlana := "senha123"
	hash, err := bcrypt.GenerateFromPassword([]byte(senhaPlana), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Erro ao gerar hash da senha: %v", err)
	}

	// 2. Montando o documento do aluno
	table_aluno.CreateAluno(Ctx, Client, hash, "unC", "20260001", "Guilherme Silva Cavalcante", []string{"MAT00131"})

	// criar um curso
	_, err = Client.Collection("cursos").Doc("CCO").Set(Ctx, map[string]interface{}{
		"codigo":  "CCO",
		"nome":    "Ciência da Computação",
		"campus":  "Darcy Ribeiro",
		"ativo":   true,
		"created": time.Now(),
	})
	if err != nil {
		log.Println("erro ao criar curso:", err)
	}

	// criar matéria
	_, err = Client.Collection("materias").Doc("MAT101").Set(Ctx, map[string]interface{}{
		"codigo":           "MAT101",
		"nome":             "Cálculo I",
		"creditos":         4,
		"cargaHoraria":     60,
		"preRequisitosIds": []string{},
		"coRequisitosIds":  []string{},
		"equivalenciasIds": []string{},
		"ativa":            true,
	})
	if err != nil {
		log.Println("erro ao criar matéria:", err)
	}

	// criar turma
	_, err = Client.Collection("turmas").Doc("T2026-1-MAT101-01").Set(Ctx, map[string]interface{}{
		"codigoTurma":    "T2026-1-MAT101-01",
		"materiaId":      "MAT101",
		"nomeMateria":    "Cálculo I",
		"semestre":       "2026.1",
		"capacidade":     40,
		"ocupadas":       0,
		"vagasRestantes": 40,
		"status":         "aberta",
	})
	if err != nil {
		log.Println("erro ao criar turma:", err)
	}

	log.Println(" Seed finalizado")
}
