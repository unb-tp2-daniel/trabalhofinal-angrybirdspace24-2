// models/aluno.go
package models

import (
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/enums"
)

// Aluno representa o perfil de um estudante matriculado.
type Aluno struct {
	Matricula          string               `json:"matricula" firestore:"matricula"`
	Nome               string               `json:"nomeAluno" firestore:"nomeAluno"`
	CursoId            string               `json:"cursoId" firestore:"cursoId"`
	Ativo              bool                 `json:"ativo" firestore:"ativo"`
	Semestre           string               `json:"semestre" firestore:"semestre"`
	MateriasConcluidas map[string]string    `json:"materiasConcluidas" firestore:"materiasConcluidas"`
	Ira                float64              `json:"ira" firestore:"ira"`
	Prioridades        map[string]string    `json:"prioridades" firestore:"prioridades"`
	NivelAcademico     enums.NivelAcademico `json:"nivelAcademico" firestore:"nivelAcademico"`
}
