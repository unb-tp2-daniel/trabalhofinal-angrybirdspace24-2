// models/departamento.go
package models

import "time"

// Departamento representa uma unidade acadêmica ou instituto.
type Departamento struct {
	DepartamentoId   string    `json:"departamentoId" firestore:"departamentoId"`
	DepartamentoNome string    `json:"departamentoNome" firestore:"departamentoNome"`
	CoordenadorId    string    `json:"coordenadorId" firestore:"coordenadorId"`
	Created          time.Time `json:"created" firestore:"created"` // Data de criação do registro
}
