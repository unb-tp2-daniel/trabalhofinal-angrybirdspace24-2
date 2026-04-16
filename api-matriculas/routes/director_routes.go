package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func DirectorRoutes() {
	//http.HandleFunc("/director", handlers.TesteHandler)
	http.HandleFunc("/director/subject/add", handlers.TesteHandler)
	http.HandleFunc("/director/subject/delete", handlers.TesteHandler)
	http.HandleFunc("/director/subject/update", handlers.TesteHandler)
}
