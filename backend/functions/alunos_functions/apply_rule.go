package alunos_functions

import (
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/enums"
)

func ApplyRules(aluno models.Aluno, materia models.Materia, curso models.Curso) float64 {
    var pontuacao float64 = 0
    var semestreAtual = "261"

    if aluno.Semestre == semestreAtual && aluno.NivelAcademico == enums.Graduacao { // calouros: prioridade máxima
        return 100_000_000 + aluno.Ira
    }

    if curso.Obrigatorias[materia.CodigoMateria] { // materia obrigatoria
        pontuacao += 10_000_000

        // desempate 1: daces
        if aluno.EhDaces {
            pontuacao += 1_000_000
        }

        // desempate 2: provavel formando
        if (aluno.Horas + materia.CargaHoraria) >= curso.TotalHoras {
            pontuacao += 100_000
        }

		// desempate 3: falta pouco pra formar
        pontuacao += (float64(aluno.Horas) / float64(curso.TotalHoras)) * 10_000 // nota de 0 a 10_000

        // desempate 4: aderencia ao curso
        // pontuacao += buscarPontosFluxo(aluno, materia)

        // ultimo criterio: ira
        pontuacao += aluno.Ira
        return pontuacao
    } else if curso.Optativas[materia.CodigoMateria] {
		pontuacao += 100_000

		// segundo o SAA, optativas entram em critérios gerais de desempate
    	pontuacao += (float64(aluno.Horas) / float64(curso.TotalHoras)) * 100
    	pontuacao += aluno.Ira
	} else { // modulo livre
		pontuacao += 10_000

		pontuacao += (float64(aluno.Horas) / float64(curso.TotalHoras)) * 100
    	pontuacao += aluno.Ira
	}

    return pontuacao
}