package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func TeacherRoutes() {
	http.HandleFunc("/teacher/subject/view", handlers.TesteHandler)
}
