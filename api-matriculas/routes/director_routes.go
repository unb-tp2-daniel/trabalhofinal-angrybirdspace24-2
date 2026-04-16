package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func DirectorRoutes() {
	http.HandleFunc("/director", handlers.TesteHandler)
	http.HandleFunc("/director/add/subject", handlers.TesteHandler)
	http.HandleFunc("/director/delete/subject", handlers.TesteHandler)
	http.HandleFunc("/director/update/subject", handlers.TesteHandler)
}
