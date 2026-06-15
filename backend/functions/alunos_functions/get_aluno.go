// functions/admin_functions/get_rules.go
package alunos_functions

import (
	"encoding/json"
	"net/http"
	"log"
	// Importando as nossas pastas isoladas
	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func GetAlunoByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	aluno, err := read.GetAlunoById(database.Ctx, database.Client, id)

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

func GetAlunoComCursoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	aluno, err := read.GetAlunoById(
		database.Ctx,
		database.Client,
		id,
	)

	if err != nil {
		log.Printf("Erro ao buscar aluno: %v", err)
		http.Error(w, "Aluno não encontrado", http.StatusInternalServerError)
		return
	}

	curso, err := read.GetCursoById(
		database.Ctx,
		database.Client,
		aluno.CursoId,
	)

	if err != nil {
		log.Printf("Erro ao buscar curso: %v", err)
		http.Error(w, "Curso não encontrado", http.StatusInternalServerError)
		return
	}

	resposta := models.AlunoDetalhado{
		Aluno: aluno,
		Curso: curso,
	}
	
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resposta); err != nil {
		log.Printf("Erro ao serializar resposta: %v", err)
		http.Error(w, "Erro interno", http.StatusInternalServerError)
	}
}