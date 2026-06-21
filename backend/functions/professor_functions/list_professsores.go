package professor_functions

import (
	"encoding/json"
	"log"
	"net/http"

	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	professorDB "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read/professor"
)

func ListProfessoresHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	// verifica se é GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido. Use GET.", http.StatusMethodNotAllowed)
		return
	}

	//chama a função do banco de dados
	professores, err := professorDB.GetAllProfessores(database.Ctx, database.Client)

	if err != nil {
		log.Printf("Erro ao buscar professores no banco: %v", err)
		http.Error(w, "Erro interno ao ler professores", http.StatusInternalServerError)
		return
	}

	//json pro front
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(professores); err != nil {
		log.Printf("Erro ao retornar JSON: %v", err)
		http.Error(w, "Erro na resposta", http.StatusInternalServerError)
	}
}
