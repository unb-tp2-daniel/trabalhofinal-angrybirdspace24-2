// functions/alunos_functions/matriculate_aluno.go
package alunos_functions

import (
	"encoding/json"
	"log"
	"net/http"

	// Importando as nossas pastas isoladas
	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/create"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

// MatriculateAlunoHandler lida exclusivamente com a requisição da internet
func MatriculateAlunoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}
	var matricula models.Matricula
	if err := json.NewDecoder(r.Body).Decode(&matricula); err != nil {
		log.Printf("Erro ao decodificar JSON da matrícula: %v", err)
		http.Error(w, "Formato de dados inválido", http.StatusBadRequest)
		return
	}

	if disponivel, err := read.TurmaIsAvailable(database.Ctx, database.Client, matricula.TurmaId); err != nil {
		http.Error(w, "Erro ao verificar disponibilidade da turma", http.StatusInternalServerError)
		return
	} else if !disponivel {
		http.Error(w, "Turma não disponível para matrícula", http.StatusConflict)
		return
	}

	err := create.CreateMatricula(database.Ctx, database.Client, matricula)
	if err != nil {
		log.Printf("Erro ao salvar matrícula no banco: %v", err)
		http.Error(w, "Erro interno ao salvar a matrícula", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Matrícula cadastrada com sucesso no banco de dados!"))
}
