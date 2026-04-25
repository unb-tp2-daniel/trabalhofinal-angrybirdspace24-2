package routes

import (
	"net/http"
	"trabalho/api-matriculas/handlers"
	"trabalho/api-matriculas/middleware"
)

func AuthRoutes() {
	http.HandleFunc("/auth", middleware.RequireInstitutionKey(handlers.GenerateTokenHandler))
}
