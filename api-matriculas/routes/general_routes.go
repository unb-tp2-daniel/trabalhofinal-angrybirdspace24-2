package routes

import (
	"api-matriculas/handlers"
	"api-matriculas/middleware"
	"net/http"
)

func GeneralRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Online. " + r.Method + " recebido."))
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	//**********************************************

	// curl -H "institutional_key: KEY" localhost:8080/teste_token
	http.HandleFunc("/teste_token", middleware.RequireInstitutionKey(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("funcionou"))
	}))

	//************************************************

	http.HandleFunc("/time", handlers.TesteHandler)
	http.HandleFunc("/status", handlers.TesteHandler)
	http.HandleFunc("/bdcapacity", handlers.TesteHandler)
	http.HandleFunc("/login", middleware.RequireInstitutionKey(handlers.LoginHandler))
}
