// functions/general_functions/server_status.go
package general_functions

import (
	"net/http"
	// Importando as nossas pastas isoladas
)

// GetServerStatusHandler lida exclusivamente com a requisição da internet
func GetServerStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Servidor online!"))
}
