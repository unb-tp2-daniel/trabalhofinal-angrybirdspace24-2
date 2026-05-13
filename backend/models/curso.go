// models/curso.go
package models

import "time"

// Curso representa um curso oferecido pela universidade.
type Curso struct {
	CursoId         string    `json:"cursoid" firestore:"cursoid"`
	CordenadorId    string    `json:"cordenadorid" firestore:"cordenadorid"`
	Nome            string    `json:"nome" firestore:"nome"`
	Campus          string    `json:"campus" firestore:"campus"`
	Ativo           bool      `json:"ativo" firestore:"ativo"`
	CargaHorariaMax int       `json:"cargaHorariaMax" firestore:"cargaHorariaMax"`
	Created         time.Time `json:"created" firestore:"created"` // Guarda a data/hora exata da criação
}
