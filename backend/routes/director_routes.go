package routes

import (
	"net/http"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/handlers"
)

func DirectorRoutes() {
	http.HandleFunc("/director", handlers.PingHandler)
	http.HandleFunc("/director/subject/add", handlers.PingHandler)
	http.HandleFunc("/director/subject/delete", handlers.PingHandler)
	http.HandleFunc("/director/subject/update", handlers.PingHandler)

	http.HandleFunc("/director/class/add", handlers.PingHandler)
	http.HandleFunc("/director/class/delete", handlers.PingHandler)
	http.HandleFunc("/director/class/update", handlers.PingHandler)

	/*

		Diretor/Coordenador:
			cadastrar materias
			permissão pra adicionar/remover aluno
			criar novas turmas
			aumentar tamanho das turmas

	*/
}
