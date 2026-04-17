package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func TeacherRoutes() {
	http.HandleFunc("/teacher/subject/show", handlers.TesteHandler)
}
