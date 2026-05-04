package database

import (
	"context"

	"time"

	"log"

	"cloud.google.com/go/firestore"
)

// Client vai guardar a conexão com o banco para a API usar

func CreateCurso(Ctx context.Context, Client *firestore.Client, codigo string, nome string, campus string, ativo bool) map[string]interface{} {

	// criar um curso
	var err error
	curso := map[string]interface{}{
		"codigo":  codigo,
		"nome":    nome,
		"campus":  campus,
		"ativo":   ativo,
		"created": time.Now(),
	}
	_, err = Client.Collection("cursos").Doc(codigo).Set(Ctx, curso)
	if err != nil {
		log.Println("erro ao criar curso:", err)
	}else{
		log.Println("curso criado com sucesso!")
	}
	return curso

}
