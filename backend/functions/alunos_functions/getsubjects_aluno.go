// functions/alunos_functions/getsubjects_aluno.go
package alunos_functions

import (
	"net/http"
	// Importando as nossas pastas isoladas
)

// GetSubjectsAlunoHandler lida exclusivamente com a requisição da internet
func GetSubjectsAlunoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Matrícula cadastrada com sucesso no banco de dados!"))
}
