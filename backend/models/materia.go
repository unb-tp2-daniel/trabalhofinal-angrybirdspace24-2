package models

import "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/enums"

// Requisito representa um grupo de disciplinas (Condição "E")
type Requisito struct {
	Disciplinas []string `json:"disciplinas" firestore:"disciplinas"`
}

// Materia representa uma disciplina oferecida nos departamentos.
type Materia struct {
	CodigoMateria  string               `json:"codigo" firestore:"codigo"`
	PreRequisitos  []Requisito          `json:"preRequisitos" firestore:"preRequisitos"` // Modificado
	DepartamentoId string               `json:"departamentoId" firestore:"departamentoId"`
	CoRequisitos   []Requisito          `json:"coRequisitos" firestore:"coRequisitos"` // Modificado
	CargaHoraria   int                  `json:"cargaHoraria" firestore:"cargaHoraria"`
	Equivalencias  []string             `json:"equivalencias" firestore:"equivalencias"`
	Conteudo       string               `json:"conteudo" firestore:"conteudo"`
	NivelAcademico enums.NivelAcademico `json:"nivelAcademico" firestore:"nivelAcademico"`
}
