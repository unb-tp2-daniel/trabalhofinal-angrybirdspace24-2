package turmas_functions

import (
	"encoding/json"
	"net/http"

	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func TurmaIsAvailable(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Use POST", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		ID string `json:"alunoId"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	doc, err := database.Client.
		Collection("turmas").
		Doc(data.ID).
		Get(database.Ctx)

	if err != nil {
		http.Error(w, "Turma não encontrada", http.StatusNotFound)
		return
	}

	var turma models.Turma

	err = doc.DataTo(&turma)

	if err != nil {
		http.Error(w, "Erro ao converter dados", http.StatusInternalServerError)
		return
	}

	disponivel := turma.VagasOcupadas < turma.VagasTotais

	json.NewEncoder(w).Encode(map[string]bool{
		"disponivel": disponivel,
	})
}
