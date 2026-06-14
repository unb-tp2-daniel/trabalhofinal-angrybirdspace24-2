// models/matricula.go
package models

import "time"

// Matricula representa a relação entre um aluno e uma turma.
type Matricula struct {
	AlunoId         string            `json:"alunoId" firestore:"alunoId"`
	TurmaId         string            `json:"turmaId" firestore:"turmaId"`
	MateriaId       string			  `json:"materiaId" firestore:"materiaId"`
	Status          bool              `json:"status" firestore:"status"`
	DataSolicitacao time.Time         `json:"dataSolicitacao" firestore:"dataSolicitacao"`
	Prioridades     map[string]string `json:"prioridades" firestore:"prioridades"`
	Semestre        string            `json:"semestre" firestore:"semestre"`
	PrioridadeNota  float64			  `json:"prioridadenota" firestore:"prioridadenota"`
}
