package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func DirectorRoutes() {
	http.HandleFunc("/director", handlers.TesteHandler)
	http.HandleFunc("/director/subject/add", handlers.TesteHandler)
	http.HandleFunc("/director/subject/delete", handlers.TesteHandler)
	http.HandleFunc("/director/subject/update", handlers.TesteHandler)

	http.HandleFunc("/director/class/add", handlers.TesteHandler)
	http.HandleFunc("/director/class/delete", handlers.TesteHandler)
	http.HandleFunc("/director/class/update", handlers.TesteHandler)

	/*
	
	Diretor/Coordenador: 
		cadastrar materias
		permissão pra adicionar/remover aluno
		criar novas turmas
		aumentar tamanho das turmas

	*/
}