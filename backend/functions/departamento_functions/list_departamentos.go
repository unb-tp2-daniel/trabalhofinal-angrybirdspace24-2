package departamento_functions

import (
	"encoding/json"
	"log"
	"net/http"

	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read"
)

// ListDepartamentosHandler atende a requisição
func ListDepartamentosHandler(w http.ResponseWriter, r *http.Request) {
	// Tornando o acesso visível para o front
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// verifica se é GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido. Use GET.", http.StatusMethodNotAllowed)
		return
	}

	deptos, err := read.GetAllDepartamentosResumo(database.Ctx, database.Client)

	if err != nil {
		log.Printf("Erro ao buscar departamentos no banco: %v", err)
		http.Error(w, "Erro interno ao ler departamentos", http.StatusInternalServerError)
		return
	}

	// Transforma a lista de structs no JSON que vai pro front e script
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(deptos); err != nil {
		log.Printf("Erro ao retornar JSON: %v", err)
		http.Error(w, "Erro na resposta", http.StatusInternalServerError)
	}
}
