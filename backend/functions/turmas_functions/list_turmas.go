package turmas_functions

import (
	"encoding/json"
	"log"
	"net/http"

	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	turmaDB "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read/turma"
)

func ListTurmasHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	// verifica se é GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido. Use GET.", http.StatusMethodNotAllowed)
		return
	}

	//chama a função do banco de dados
	turmas, err := turmaDB.GetTurmas(database.Ctx, database.Client, 5)

	if err != nil {
		log.Printf("Erro ao buscar turmas no banco: %v", err)
		http.Error(w, "Erro interno ao ler turmas", http.StatusInternalServerError)
		return
	}

	//json pro front
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(turmas); err != nil {
		log.Printf("Erro ao retornar JSON: %v", err)
		http.Error(w, "Erro na resposta", http.StatusInternalServerError)
	}
}
