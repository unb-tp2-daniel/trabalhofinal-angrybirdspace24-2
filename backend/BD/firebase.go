package database

import (
	"context"
	"log"

	"os"

	"cloud.google.com/go/firestore"
	//tables "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/BD/tables"
	//"golang.org/x/crypto/bcrypt"
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
