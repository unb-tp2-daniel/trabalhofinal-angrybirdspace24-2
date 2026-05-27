// models/materia.go
package models

import "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/enums"

// Materia representa uma disciplina oferecida nos departamentos.
type Materia struct {
	CodigoMateria  string               `json:"codigo" firestore:"codigo"`
	PreRequisitos  [][]string           `json:"preRequisitos" firestore:"preRequisitos"`
	DepartamentoId string               `json:"departamentoId" firestore:"departamentoId"`
	CoRequisitos   [][]string           `json:"coRequisitos" firestore:"coRequisitos"`
	CargaHoraria   int                  `json:"cargaHoraria" firestore:"cargaHoraria"`
	Equivalencias  []string             `json:"equivalencias" firestore:"equivalencias"` // Alterado para slice
	Conteudo       string               `json:"conteudo" firestore:"conteudo"`
	NivelAcademico enums.NivelAcademico `json:"nivelAcademico" firestore:"nivelAcademico"`
}
