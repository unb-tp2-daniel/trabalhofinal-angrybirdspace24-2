package database

import (
	"context"

	"time"

	"log"

	"cloud.google.com/go/firestore"
)

// Client vai guardar a conexão com o banco para a API usar

func CreateCurso(Ctx context.Context, Client *firestore.Client, cursoid string, nome string, campus string, ativo bool, cargaHorariaMax int, cordenadorid string) map[string]interface{} {

	// criar um curso
	var err error
	curso := map[string]interface{}{
		"cursoid":  cursoid,
		"cordenadorid": cordenadorid,
		"nome":    nome,
		"campus":  campus,
		"ativo":   ativo,
		"cargaHorariaMax": cargaHorariaMax,
		"created": time.Now(),
	}
	_, err = Client.Collection("cursos").Doc(cursoid).Set(Ctx, curso)
	if err != nil {
		log.Println("erro ao criar curso:", err)
	}else{
		log.Println("curso criado com sucesso!")
	}
	return curso

}
