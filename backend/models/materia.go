// models/materia.go
package models

import "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/enums"

// Materia representa uma disciplina oferecida nos departamentos.
type Materia struct {
	CodigoMateria  string               `json:"codigo" firestore:"codigo"`
	PreRequisitos  map[string]string    `json:"preRequisitos" firestore:"preRequisitos"`
	DepartamentoId string               `json:"departamentoId" firestore:"departamentoId"`
	CoRequisitos   map[string]string    `json:"coRequisitos" firestore:"coRequisitos"`
	CargaHoraria   int                  `json:"cargaHoraria" firestore:"cargaHoraria"`
	Equivalencias  map[string]string    `json:"equivalencias" firestore:"equivalencias"`
	Conteudo       string               `json:"conteudo" firestore:"conteudo"`
	NivelAcademico enums.NivelAcademico `json:"nivelAcademico" firestore:"nivelAcademico"`
	Prioridades    map[string]string    `json:"prioridades" firestore:"prioridades"`
}
