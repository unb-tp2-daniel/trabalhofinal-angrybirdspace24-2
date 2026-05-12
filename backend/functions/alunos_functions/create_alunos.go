// functions/alunos_functions/create_aluno.go
package alunos_functions

import (
	"encoding/json"
	"log"
	"net/http"

	// Importando as nossas pastas isoladas
	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/BD"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/BD/create"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/models"
)

// CreateAlunoHandler lida exclusivamente com a requisição da internet
func CreateAlunoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}
	var novoAluno models.Aluno
	if err := json.NewDecoder(r.Body).Decode(&novoAluno); err != nil {
		log.Printf("Erro ao decodificar JSON do aluno: %v", err)
		http.Error(w, "Formato de dados inválido", http.StatusBadRequest)
		return
	}
	err := create.CreateAluno(database.Ctx, database.Client, novoAluno)
	if err != nil {
		log.Printf("Erro ao salvar aluno no banco: %v", err)
		http.Error(w, "Erro interno ao salvar o aluno", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Aluno criado com sucesso no banco de dados!"))
}
