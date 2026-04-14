package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/teste", handlers.UsersHandler)
}
