package database

import (
	"context"

	"time"

	"log"

	"cloud.google.com/go/firestore"
)

// Client vai guardar a conexão com o banco para a API usar

func CreateCoordenador(Ctx context.Context, 
	Client *firestore.Client,
	coordenadorid string,
	coordenadorNome string,
	departamentoid string,
	) map[string]interface{} {

	// criar um curso
	var err error
	coordenador := map[string]interface{}{
		"coordenadorId": coordenadorid,
		"departamentoId": departamentoid,
		"coordenadorNome": coordenadorNome,
		"created": time.Now(),
	}
	_, err = Client.Collection("coordenadores").Doc(coordenadorid).Set(Ctx, coordenador)
	if err != nil {
		log.Println("erro ao criar coordenador:", err)
	}else{
		log.Println("coordenador criado com sucesso!")
	}
	return coordenador

}
