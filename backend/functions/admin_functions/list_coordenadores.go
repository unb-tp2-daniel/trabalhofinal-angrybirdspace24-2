package admin_functions

import (
	"encoding/json"
	"log"
	"net/http"

	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	alunoDB "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read/coordenador"
)

func ListCoordenadoresHandler(w http.ResponseWriter, r *http.Request) {
	// verifica se é GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido. Use GET.", http.StatusMethodNotAllowed)
		return
	}

	//chama a função do banco de dados
	coordenadores, err := alunoDB.GetAllCoordenadores(database.Ctx, database.Client)

	if err != nil {
		log.Printf("Erro ao buscar coordenadores no banco: %v", err)
		http.Error(w, "Erro interno ao ler coordenadores", http.StatusInternalServerError)
		return
	}

	//json pro front
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(coordenadores); err != nil {
		log.Printf("Erro ao retornar JSON: %v", err)
		http.Error(w, "Erro na resposta", http.StatusInternalServerError)
	}
}