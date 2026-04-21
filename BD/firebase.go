package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// Client vai guardar a conexão com o banco para a API usar
var Client *firestore.Client

func InitDB() {
	ctx := context.Background()

	opt := option.WithAuthCredentialsFile(option.ServiceAccount, "../serviceAccountKey.json")

	//Conecta diretamente no banco matriculas242
	var err error
	Client, err = firestore.NewClientWithDatabase(ctx, "matriculas242", "matriculas242", opt)

	if err != nil {
		log.Fatalf("Erro ao conectar no Firestore: %v", err)
	}

	log.Println("✅ BD online")
}
