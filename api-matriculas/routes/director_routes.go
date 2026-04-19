package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func DirectorRoutes() {
	http.HandleFunc("/director", handlers.TestHandler)
	http.HandleFunc("/director/subject/add", handlers.TestHandler)
	http.HandleFunc("/director/subject/delete", handlers.TestHandler)
	http.HandleFunc("/director/subject/update", handlers.TestHandler)

	http.HandleFunc("/director/class/add", handlers.TestHandler)
	http.HandleFunc("/director/class/delete", handlers.TestHandler)
	http.HandleFunc("/director/class/update", handlers.TestHandler)

	/*
	
	Diretor/Coordenador: 
		cadastrar materias
		permissão pra adicionar/remover aluno
		criar novas turmas
		aumentar tamanho das turmas

	*/
}