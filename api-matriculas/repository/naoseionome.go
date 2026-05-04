package repository

import (
	"context"
	"fmt"

	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/BD"

	"google.golang.org/api/iterator"
)

func FindInstitutionByID(institutionalKey string) (string, error) {
	ctx := context.Background()

	iter := database.Client.Collection("instituicoes").Where("chave_acesso", "==", institutionalKey).Limit(1).Documents(ctx)
	// Tenta ler o primeiro resultado
	doc, err := iter.Next()

	// FB retorna iterator.Done qnd a query falha
	if err == iterator.Done {
		return "", fmt.Errorf("chave inválida")
	}
	// erro externoo
	if err != nil {
		return "", fmt.Errorf("erro interno ao buscar instituição: %v", err)
	}
	//se a query der certo vai ser "UnB"
	return doc.Ref.ID, nil
}
