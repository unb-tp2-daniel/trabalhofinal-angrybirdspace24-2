package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func DiretorRoutes() {
	http.HandleFunc("/alunos", handlers.TesteHandler)
	http.HandleFunc("/alunos/criar", handlers.TesteHandler)
}
