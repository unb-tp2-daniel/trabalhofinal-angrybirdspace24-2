// functions/admin_functions/get_rules.go
package alunos_functions

import (
	"encoding/json"
	"net/http"
	"log"
	// Importando as nossas pastas isoladas
	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read"
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
