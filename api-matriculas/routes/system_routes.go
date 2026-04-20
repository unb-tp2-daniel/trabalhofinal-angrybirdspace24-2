package routes

import (
	"api-matriculas/handlers"
	"api-matriculas/middleware"
	"net/http"
)

func AuthRoutes() {
	http.HandleFunc("/auth", middleware.RequireInstitutionKey(handlers.GenerateTokenHandler))
}
