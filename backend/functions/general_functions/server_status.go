package general_functions

import (
	"net/http"
)

func GetServerStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Servidor online!"))
}
