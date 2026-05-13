// models/coordenador.go
package models

// Coordenador representa o responsável por um departamento específico.
type Coordenador struct {
	CoordenadorId   string `json:"coordenadorId" firestore:"coordenadorId"`
	DepartamentoId  string `json:"departamentoId" firestore:"departamentoId"`
	CoordenadorNome string `json:"coordenadorNome" firestore:"coordenadorNome"`
}
