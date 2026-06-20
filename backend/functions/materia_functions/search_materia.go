package materia_functions

import (
	"encoding/json"
	"log"
	"net/http"

	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read"
	
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SearchMateriaHandler busca uma matéria específica pelo ID passado na URL
func SearchMateriaHandler(w http.ResponseWriter, r *http.Request) {
	// Tornando o acesso visível para o front-end
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Verifica se o método é GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido. Use GET.", http.StatusMethodNotAllowed)
		return
	}

	// Extrai o ID da query string da URL 
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Parâmetro 'id' é obrigatório", http.StatusBadRequest)
		return
	}

	// Chama a função do banco de dados (getById)
	materia, err := read.GetMateriaById(database.Ctx, database.Client, id)

	if err != nil {
		// Verifica se o erro ocorreu porque o documento não existe no Firestore
		if status.Code(err) == codes.NotFound {
			http.Error(w, "Matéria não encontrada", http.StatusNotFound)
			return
		}
		
		// Caso seja outro tipo de erro interno
		log.Printf("Erro ao buscar materia %s no banco: %v", id, err)
		http.Error(w, "Erro interno ao ler matéria", http.StatusInternalServerError)
		return
	}

	// Retorna o JSON para o front-end
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(materia); err != nil {
		log.Printf("Erro ao retornar JSON: %v", err)
		http.Error(w, "Erro na resposta", http.StatusInternalServerError)
	}
}
