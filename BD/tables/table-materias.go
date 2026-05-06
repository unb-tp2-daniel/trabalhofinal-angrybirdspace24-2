package database

import (
	"context"

	"log"

	"cloud.google.com/go/firestore"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/BD/enums"
)

// Client vai guardar a conexão com o banco para a API usar

func CreateMateria(Ctx context.Context, Client *firestore.Client,
	codigoMateria string,
	preRequisitos map[string]string,
	departamentoId string,
	coRequisitos map[string]string,
	cargaHorario int,
	equivalencias map[string]string,
	conteudo string,
	nivelAcademico enums.NivelAcademico,
	prioridades map[string]string) map[string]interface{} {

	// criar um curso
	var err error
	materia := map[string]interface{}{
		"codigo":         codigoMateria,
		"preRequisitos":  preRequisitos,
		"departamentoId": departamentoId,
		"coRequisitos":   coRequisitos,
		"cargaHorario":   cargaHorario,
		"equivalencias":  equivalencias,
		"conteudo":       conteudo,
		"nivelAcademico": nivelAcademico,
	}
	_, err = Client.Collection("materias").Doc(codigoMateria).Set(Ctx, materia)
	if err != nil {
		log.Println("erro ao criar materia:", err)
	} else {
		log.Println("Materia criada com sucesso!")
	}
	return materia

}
