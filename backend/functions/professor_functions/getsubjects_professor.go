// functions/professor_functions/getsubjects_professor.go
package professor_functions

import (
	"net/http"
	// Importando as nossas pastas isoladas
)

// GetSubjectsProfessorHandler lida exclusivamente com a requisição da internet
func GetSubjectsProfessorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Matrícula cadastrada com sucesso no banco de dados!"))
}
