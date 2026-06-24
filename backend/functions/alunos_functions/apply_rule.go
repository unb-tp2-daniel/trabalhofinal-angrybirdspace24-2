package alunos_functions

import (
	time "time"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/enums"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func ApplyRules(aluno models.Aluno, materia models.Materia, curso models.Curso) float64 {
	var pontuacao float64 = 0

	if isCalouro(aluno) { // calouros: prioridade máxima
		return 100_000_000 + aluno.Ira
	}

	pontuacao += calcularPontosBase(materia.CodigoMateria, curso)
	pontuacao += calcularProgresso(aluno, curso, materia)
	pontuacao += aluno.Ira

	return pontuacao
}

func isCalouro(aluno models.Aluno) bool {
	ano := time.Now().Format("06")
	semestre := "1"

	if time.Now().Month() > 7 {
		semestre = "2"
	}

	semestreAtual := ano + semestre
	return aluno.Semestre == semestreAtual && aluno.NivelAcademico == enums.Graduacao
}

func calcularPontosBase(materiaId string, curso models.Curso) float64 {
	if curso.Obrigatorias[materiaId] {
		return 10_000_000
	}

	if curso.Optativas[materiaId] {
		return 100_000
	}

	return 10_000
}

func calcularProgresso(aluno models.Aluno, curso models.Curso, materia models.Materia) float64 {
	if !(curso.Obrigatorias[materia.CodigoMateria]) {
		// segundo o SAA, optativas entram em critérios gerais de desempate
		return ((float64(aluno.Horas) / float64(curso.TotalHoras)) * 100)
	}

	var bonus float64
	if aluno.Prioridades["Daces"] == "sim" {
		bonus += 1_000_000
	}

	//provavel formando
	if (aluno.Horas + materia.CargaHoraria) >= curso.TotalHoras {
		bonus += 100_000
	}

	// desempate por fluxo
	// pontuacao += buscarPontosFluxo(aluno, materia)

	bonus += (float64(aluno.Horas) / float64(curso.TotalHoras)) * 10_000 // nota de 0 a 10_000
	return bonus
}
