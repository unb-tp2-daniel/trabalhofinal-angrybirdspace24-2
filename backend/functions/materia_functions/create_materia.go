package materia_functions

import (
	"encoding/json"
	"log"
	"net/http"

	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/create" // aaqui salva
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func CreateMateriaHandler(w http.ResponseWriter, r *http.Request) {
	// verifica se o é POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido. Use POST.", http.StatusMethodNotAllowed)
		return
	}

	var novaMateria models.Materia

	// le o json do front e joga no model
	if err := json.NewDecoder(r.Body).Decode(&novaMateria); err != nil {
		log.Printf("Erro ao decodificar JSON: %v", err)
		http.Error(w, "Formato de dados inválido", http.StatusBadRequest)
		return
	}

	//cria de vdd só na pasta BD
	err := create.CreateMateria(database.Ctx, database.Client, novaMateria)

	if err != nil {
		log.Printf("Erro ao salvar no banco: %v", err)
		http.Error(w, "Erro interno ao salvar a turma", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Turma criada com sucesso no Firebase!"))
}
