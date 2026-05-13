// functions/admin_functions/get_rules.go
package admin_functions

import (
	"net/http"
	// Importando as nossas pastas isoladas
)

// GetRulesHandler lida exclusivamente com a requisição da internet
func GetRulesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Regras do sistema obtidas com sucesso!"))
}
