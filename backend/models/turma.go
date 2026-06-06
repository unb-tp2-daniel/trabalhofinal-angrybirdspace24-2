// models/turma.go
package models

// Turma representa uma classe de uma matéria ofertada.
// As tags `json` dizem como o frontend manda os dados.
// As tags `firestore` dizem como o banco de dados vai salvar as chavess.
type Turma struct {
	CodigoTurma     string            `json:"codigoTurma" firestore:"codigoTurma"`
	MateriaId       string            `json:"materiaId" firestore:"materiaId"`
	NomeMateria     string            `json:"nomeMateria" firestore:"nomeMateria"`
	Semestre        string            `json:"semestre" firestore:"semestre"`
	VagasTotais     int64             `json:"vagasTotais" firestore:"vagasTotais"`
	VagasOcupadas   int64             `json:"vagasOcupadas" firestore:"vagasOcupadas"`
	Ativo           bool              `json:"ativo" firestore:"ativo"`
	Local           string            `json:"local" firestore:"local"`
	Horario         string            `json:"horario" firestore:"horario"`
	VagasExclusivas map[string]int64  `json:"vagasExclusivas" firestore:"vagasExclusivas"`
	ProfessorId     string            `json:"professorId" firestore:"professorid"`
	ProfessorNome   string            `json:"professorNome" firestore:"professorNome"`
	Prioridades     map[string]string `json:"prioridades" firestore:"prioridades"`
}
