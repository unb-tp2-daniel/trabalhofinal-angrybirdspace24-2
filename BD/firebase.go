package database

import (
	"context"
	"log"

	"os"

	"cloud.google.com/go/firestore"
	tables "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/BD/tables"
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
	var err error

	// A nuvem injeta a variável K_SERVICE automaticamente.
	// Se ela existir, estamos no Cloud Functions!
	if os.Getenv("K_SERVICE") != "" {
		log.Println("Conectando ao Firestore pelo ambiente Cloud (IAM automático)...")
		Client, err = firestore.NewClientWithDatabase(Ctx, "matriculas242", "matriculas242")
	} else {
		log.Println("Conectando ao Firestore pelo ambiente Local (serviceAccountKey)...")
		opt := option.WithAuthCredentialsFile(option.ServiceAccount, "../serviceAccountKey.json")
		Client, err = firestore.NewClientWithDatabase(Ctx, "matriculas242", "matriculas242", opt)
	}

	if err != nil {
		log.Fatalf("Erro ao conectar no Firestore: %v", err)
	}

	log.Println("BD online")
}

// SeedBaseData cria dados iniciais no banco
func SeedBaseData() {
	log.Println("Iniciando seed do banco...")

	// Criptografa
	senhaPlana := "senha123"
	hash, err := bcrypt.GenerateFromPassword([]byte(senhaPlana), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Erro ao gerar hash da senha: %v", err)
	}

	// criando tabelas
	tables.CreateAluno(Ctx, Client, hash, "unC", "20260001", "Guilherme Silva Cavalcante", []string{"MAT00131"})

	tables.CreateTurma(Ctx, Client, hash, "T2026-1-MAT101-01", "MAT101", "Cálculo I", "2026.1", 40, 0, 40, true)

	tables.CreateCurso(Ctx, Client, "CCO", "Ciencia da computação", "Darcy Ribeiro", true)

	tables.CreateMateria(Ctx, Client, "Mat11", "matematica", 6, 60, []string{}, []string{}, []string{}, true)

}
