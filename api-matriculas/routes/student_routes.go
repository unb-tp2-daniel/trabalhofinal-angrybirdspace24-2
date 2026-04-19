package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func StudentRoutes() {
	http.HandleFunc("/student", handlers.TestHandler)
	http.HandleFunc("/student/matriculate", handlers.TestHandler)
	http.HandleFunc("/student/position", handlers.TestHandler)
	http.HandleFunc("/student/subjects", handlers.TestHandler)
	http.HandleFunc("/student/schedule", handlers.TestHandler)
}
