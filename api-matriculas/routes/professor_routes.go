package routes

import (
	"net/http"
	"trabalho/api-matriculas/handlers"
)

func TeacherRoutes() {
	http.HandleFunc("/teacher/subject/show", handlers.TestHandler)
}
