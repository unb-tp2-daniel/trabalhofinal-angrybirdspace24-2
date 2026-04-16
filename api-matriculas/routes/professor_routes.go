package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func TeacherRoutes() {
	http.HandleFunc("/teacher", handlers.TesteHandler)
	http.HandleFunc("/teacher/create", handlers.TesteHandler)
}
