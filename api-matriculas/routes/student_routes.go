package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func StudentRoutes() {
	http.HandleFunc("/student", handlers.TesteHandler)
	http.HandleFunc("/student/matriculate", handlers.TesteHandler)
	http.HandleFunc("/student/position", handlers.TesteHandler)
	http.HandleFunc("/student/subjects", handlers.TesteHandler)
	http.HandleFunc("/student/schedule", handlers.TesteHandler)

}
