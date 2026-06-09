// functions/alunos_functions/apply_rules.go
package alunos_functions

import (
	"strconv"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

// ApplyRules calcula a nota de prioridade do aluno para a matrícula normal.
//
// Fórmula: 60% do peso vem do percentual do curso concluído + 40% do IRA normalizado.
// Resultado final: float64 no intervalo [0, 100].
//
// Exemplos:
//   integralizado=80%, IRA=5.0 → 0.6*80 + 0.4*(5/5*100) = 48 + 40 = 88.0
//   integralizado=30%, IRA=4.0 → 0.6*30 + 0.4*(4/5*100) = 18 + 32 = 50.0
//   integralizado=50%, IRA=2.5 → 0.6*50 + 0.4*(2.5/5*100) = 30 + 20 = 50.0
//
// Os pesos PESO_INTEGRALIZADO e PESO_IRA devem sempre somar 1.0.

const (
	PESO_INTEGRALIZADO = 0.6
	PESO_IRA           = 0.4
	IRA_MAXIMO         = 5.0 // IRA máximo possível na UnB
)

func ApplyRules(aluno models.Aluno, materia models.Materia, curso models.Curso) float64 {
	// --- 1. Percentual do curso concluído (0–100) ---
	// Lido do campo prioridades["integralizado"], que vem como string (ex: "94")
	integralizadoStr := aluno.Prioridades["integralizado"]
	integralizado, err := strconv.ParseFloat(integralizadoStr, 64)
	if err != nil {
		integralizado = 0
	}
	integralizado = clamp(integralizado, 0, 100)

	// --- 2. IRA normalizado para escala 0–100 ---
	iraNormalizado := (aluno.Ira / IRA_MAXIMO) * 100
	iraNormalizado = clamp(iraNormalizado, 0, 100)

	// --- 3. Nota final ponderada ---
	return (PESO_INTEGRALIZADO * integralizado) + (PESO_IRA * iraNormalizado)
}

// clamp garante que v está entre min e max.
func clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}