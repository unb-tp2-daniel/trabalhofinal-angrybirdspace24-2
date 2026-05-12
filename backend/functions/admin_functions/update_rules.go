// functions/admin_functions/update_rules.go
package admin_functions

import (
	"net/http"
	// Importando as nossas pastas isoladas
)

// UpdateRulesHandler lida exclusivamente com a requisição da internet
func UpdateRulesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Regras do sistema atualizadas com sucesso!"))
}
