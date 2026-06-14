package delete

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
)

func ClearCollection(ctx context.Context, client *firestore.Client, collectionName string) error {
	if client == nil {
		return fmt.Errorf("cliente do firestore não inicializado")
	}

	collectionRef := client.Collection(collectionName)
	const batchSize = 100

	for {
		docs, err := collectionRef.Limit(batchSize).Select().Documents(ctx).GetAll()
		if err != nil {
			return fmt.Errorf("falha ao buscar documentos para limpeza: %w", err)
		}

		if len(docs) == 0 {
			break
		}

		batch := client.Batch()
		for _, doc := range docs {
			batch.Delete(doc.Ref)
		}

		_, err = batch.Commit(ctx)
		if err != nil {
			return fmt.Errorf("falha ao comitar o lote de deleção na coleção %s: %w", collectionName, err)
		}

		log.Printf("Lote de %d documentos deletado com sucesso da coleção: %s", len(docs), collectionName)
	}

	log.Printf("coleção %s foi completamente limpa", collectionName)
	return nil
}