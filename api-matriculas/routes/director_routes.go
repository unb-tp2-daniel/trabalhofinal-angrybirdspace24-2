package routes

import (
	"net/http"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/api-matriculas/handlers"
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
