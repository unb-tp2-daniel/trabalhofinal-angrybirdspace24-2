// models/professor.go
package models

// Professor representa um docente vinculado a um departamento da universidade.
type Professor struct {
	ProfessorId    string `json:"professorid" firestore:"professorid"`
	ProfessorNome  string `json:"professorNome" firestore:"professorNome"`
	DepartamentoId string `json:"departamentoid" firestore:"departamentoid"`
}
