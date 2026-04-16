package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func GeneralRoutes() {
	http.HandleFunc("/ping", handlers.TesteHandler)
	http.HandleFunc("/time", handlers.TesteHandler)
}
