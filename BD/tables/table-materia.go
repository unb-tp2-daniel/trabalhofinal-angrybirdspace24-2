package database

import (
	"context"

	"time"

	"log"

	"cloud.google.com/go/firestore"
)

// Client vai guardar a conexão com o banco para a API usar

func CreateMateria(Ctx context.Context, Client *firestore.Client, codigo string, nome string, campus string, ativo bool) map[string]interface{} {

	// criar um curso
	var err error
	materia := map[string]interface{}{
			"codigo":           "MAT101",
		"nome":             "Cálculo I",
		"creditos":         4,
		"cargaHoraria":     60,
		"preRequisitosIds": []string{},
		"coRequisitosIds":  []string{},
		"equivalenciasIds": []string{},
		"ativa":            true,

	}
	_, err = Client.Collection("materias").Doc("MAT101").Set(Ctx, materia)
	if err != nil {
		log.Println("erro ao criar curso:", err)
	}
	return curso

}


	_, err = Client.Collection("materias").Doc("MAT101").Set(Ctx, map[string]interface{}{
		"codigo":           "MAT101",
		"nome":             "Cálculo I",
		"creditos":         4,
		"cargaHoraria":     60,
		"preRequisitosIds": []string{},
		"coRequisitosIds":  []string{},
		"equivalenciasIds": []string{},
		"ativa":            true,
	})
	if err != nil {
		log.Println("erro ao criar matéria:", err)
	}
