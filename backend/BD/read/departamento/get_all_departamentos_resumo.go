package departamento

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)


type DepartamentoResumo struct {
	DepartamentoNome string `json:"departamentoNome" firestore:"departamentoNome"`
	CoordenadorId    string `json:"coordenadorId" firestore:"coordenadorId"`
	DepartamentoId   string `json:"departamentoId" firestore:"departamentoId"`
}


func GetAllDepartamentosResumo(ctx context.Context, client *firestore.Client) ([]DepartamentoResumo, error) {
	var lista []DepartamentoResumo

	iter := client.Collection("departamentos").Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break 
		}
		if err != nil {
			return nil, err 
		}

		var resumo DepartamentoResumo
		// O DataTo joga os dados do Firestore direto pra struct enxuta
		if err := doc.DataTo(&resumo); err != nil {
			return nil, err
		}

		
		lista = append(lista, resumo)
	}

	return lista, nil
}
