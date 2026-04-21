package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

// Client vai guardar a conexão com o banco para a API usar
var Client *firestore.Client

func InitDB() {
	ctx := context.Background()

	opt := option.WithCredentialsFile("../serviceAccountKey.json")

	config := &firebase.Config{
		ProjectID: "matriculas242", //id no terminal do firebase
	}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("Erro ao inicializar Firebase: %v", err)
	}

	Client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Erro ao conectar no Firestore: %v", err)
	}

	log.Println("BD online")
}
