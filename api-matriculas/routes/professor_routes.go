package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func ProfessorRoutes() {
	http.HandleFunc("/professor", handlers.TesteHandler)
	http.HandleFunc("/professor/criar", handlers.TesteHandler)
}
