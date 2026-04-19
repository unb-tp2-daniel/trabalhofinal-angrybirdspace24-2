package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func GeneralRoutes() {
	http.HandleFunc("/", handlers.TestHandler)
	http.HandleFunc("/ping", handlers.PingHandler)
	http.HandleFunc("/time", handlers.TimeHandler)
	http.HandleFunc("/status", handlers.StatusHandler)

	/*// curl -H "institutional_key: KEY" localhost:8080/TestHandlere_token
	http.HandleFunc("/TestHandlere_token", middleware.RequireInstitutionKey(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("funcionou"))
	})) to do
	http.HandleFunc("/bdcapacity", handlers.TestHandler) To do*/

	//http.HandleFunc("/login", middleware.RequireInstitutionKey(handlers.LoginHandler))
}
