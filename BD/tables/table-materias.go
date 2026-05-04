package database

import (
	"context"

	"log"

	"cloud.google.com/go/firestore"
)

// Client vai guardar a conexão com o banco para a API usar

func CreateMateria(Ctx context.Context, Client *firestore.Client, codigo string, nome string, creditos int, cargaHoraria int, preRequisitosIds []string, coRequisitosIds []string, equivalenciasIds []string, ativo bool) map[string]interface{} {

	// criar um curso
	var err error
	materia := map[string]interface{}{
		"codigo":           codigo,
		"nome":             nome,
		"creditos":         creditos,
		"cargaHoraria":     cargaHoraria,
		"preRequisitosIds": preRequisitosIds,
		"coRequisitosIds":  coRequisitosIds,
		"equivalenciasIds": equivalenciasIds,
		"ativa":            ativo,
	}
	_, err = Client.Collection("materias").Doc(codigo).Set(Ctx, materia)
	if err != nil {
		log.Println("erro ao criar materia:", err)
	}else{
		log.Println("Materia criada com sucesso!",)
	}
	return materia

}
