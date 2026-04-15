package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func AlunoRoutes() {
	http.HandleFunc("/alunos", handlers.TesteHandler)
	http.HandleFunc("/alunos/criar", handlers.TesteHandler)
}
