package read

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// O Firestore vai ler o documento e preencher só o que bater com essas tags.
type DepartamentoResumo struct {
	DepartamentoNome string `json:"departamentoNome" firestore:"departamentoNome"`
	CoordenadorId    string `json:"coordenadorId" firestore:"coordenadorId"`
	DepartamentoId   string `json:"departamentoId" firestore:"departamentoId"`
}

// GetAllDepartamentosResumo busca todos os deptos e retorna só o nome e a ID do coordenador
func GetAllDepartamentosResumo(ctx context.Context, client *firestore.Client) ([]DepartamentoResumo, error) {
	var lista []DepartamentoResumo

	iter := client.Collection("departamentos").Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break // Fim da lista de documentos
		}
		if err != nil {
			return nil, err // Retorna o erro pro Handler lidar
		}

		var resumo DepartamentoResumo
		// O DataTo joga os dados do Firestore direto pra struct enxuta
		if err := doc.DataTo(&resumo); err != nil {
			return nil, err
		}

		// Adiciona na lista
		lista = append(lista, resumo)
	}

	return lista, nil
}
