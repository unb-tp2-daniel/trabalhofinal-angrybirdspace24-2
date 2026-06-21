// functions/alunos_functions/getposition_aluno.go
package alunos_functions

import (
	"net/http"
)

func GetPositionAlunoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Matrícula cadastrada com sucesso no banco de dados!"))
}
