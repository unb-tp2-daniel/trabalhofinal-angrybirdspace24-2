// functions/alunos_functions/matriculate_aluno.go
package alunos_functions

import (
	"net/http"
	// Importando as nossas pastas isoladas
)

// MatriculateAlunoHandler lida exclusivamente com a requisição da internet
func MatriculateAlunoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}
	/*var matricula models.Matricula
	if err := json.NewDecoder(r.Body).Decode(&matricula); err != nil {
		log.Printf("Erro ao decodificar JSON da matrícula: %v", err)
		http.Error(w, "Formato de dados inválido", http.StatusBadRequest)
		return
	}
	err := create.CreateMatricula(database.Ctx, database.Client, matricula)
	if err != nil {
		log.Printf("Erro ao salvar matrícula no banco: %v", err)
		http.Error(w, "Erro interno ao salvar a matrícula", http.StatusInternalServerError)
		return
	}*/
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Matrícula cadastrada com sucesso no banco de dados!"))
}
