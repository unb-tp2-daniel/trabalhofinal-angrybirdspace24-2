package routes

import (
	"api-matriculas/handlers"
	"api-matriculas/middleware"
	"net/http"
)

func AuthRoutes() {
	http.HandleFunc("/auth/student", middleware.RequireInstitutionKey(handlers.GenerateStudentToken))
	http.HandleFunc("/auth/teacher", middleware.RequireInstitutionKey(handlers.GenerateTeacherToken))
}
