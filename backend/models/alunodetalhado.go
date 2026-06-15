// models/aluno.go
package models

// Aluno representa o perfil de um estudante matriculado.
type AlunoDetalhado struct {
	*Aluno
	Curso *Curso `json:"curso"`
}
