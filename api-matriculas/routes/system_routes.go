package routes

import (
	"api-matriculas/handlers"
	"api-matriculas/middleware"
	"net/http"
)

func AuthRoutes() {
	http.HandleFunc("/auth/student", middleware.RequireInstitutionKey(handlers.GenerateTokenHandler))
	http.HandleFunc("/auth/professor", middleware.RequireInstitutionKey(handlers.GenerateTokenHandler))
	http.HandleFunc("/auth/director", middleware.RequireInstitutionKey(handlers.GenerateTokenHandler))
}
