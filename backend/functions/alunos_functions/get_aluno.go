// functions/admin_functions/get_rules.go
package alunos_functions

import (
	"encoding/json"
	"log"
	"net/http"

	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	alunoDB "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read/aluno"
	cursoDB "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read/curso"
	materiaDB "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read/materia"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func GetAlunoByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	aluno, err := alunoDB.GetAlunoById(database.Ctx, database.Client, id)

	if err != nil {
		log.Printf("Erro ao buscar aluno no banco: %v", err)
		http.Error(w, "Erro interno ao ler aluno", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(aluno); err != nil {
		log.Printf("Erro ao retornar JSON: %v", err)
		http.Error(w, "Erro na resposta", http.StatusInternalServerError)
	}

}

func GetCHHandler(w http.ResponseWriter, r *http.Request) {
	//Tornando o acesso visível para o front
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// verifica se é GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido. Use GET.", http.StatusMethodNotAllowed)
		return
	}

	idAluno := r.URL.Query().Get("idAluno")
	aluno, err := alunoDB.GetAlunoById(database.Ctx, database.Client, idAluno)
	curso, err := cursoDB.GetCursoById(database.Ctx, database.Client, aluno.CursoId)

	var materiasObrigatorias []*models.Materia
	var materiasOptativas []*models.Materia

	var chObrigatorias int
	var chOptativas int

	for id_materia, _ := range curso.Obrigatorias {
		materia, _ := materiaDB.GetMateriaById(database.Ctx, database.Client, id_materia)
		materiasObrigatorias = append(materiasObrigatorias, materia)
		chObrigatorias += materia.CargaHoraria
	}
	for id_materia, _ := range curso.Optativas {
		materia, _ := materiaDB.GetMateriaById(database.Ctx, database.Client, id_materia)
		materiasOptativas = append(materiasOptativas, materia)
		chOptativas += materia.CargaHoraria
	}

	var chObrigatoriaPendente int
	var chOptativaPendente int

	for _, m := range materiasObrigatorias {
		_, existe := aluno.MateriasConcluidas[m.CodigoMateria]
		if !existe {
			chObrigatoriaPendente += m.CargaHoraria
		}
	}
	for _, m := range materiasOptativas {
		_, existe := aluno.MateriasConcluidas[m.CodigoMateria]
		if !existe {
			chOptativaPendente += m.CargaHoraria
		}
	}

	if err != nil {
		log.Printf("Erro ao buscar materias no banco: %v", err)
		http.Error(w, "Erro interno ao ler materias", http.StatusInternalServerError)
		return
	}

	//json pro front
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]any{
		"chObrigatorias":         chObrigatorias,
		"chOptativas":            chOptativas,
		"chObrigatoriasPendente": chObrigatoriaPendente,
		"chOptativasPendente":    chOptativaPendente,
	}); err != nil {
		log.Printf("Erro ao retornar JSON: %v", err)
		http.Error(w, "Erro na resposta", http.StatusInternalServerError)
	}
}
