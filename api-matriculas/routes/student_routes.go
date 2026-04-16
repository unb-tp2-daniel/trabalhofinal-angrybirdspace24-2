package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func StudentRoutes() {
	http.HandleFunc("/student", handlers.TesteHandler)
	http.HandleFunc("/student/create", handlers.TesteHandler)
}
