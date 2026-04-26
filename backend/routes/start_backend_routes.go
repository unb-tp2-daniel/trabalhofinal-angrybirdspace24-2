package routes

import (
	"net/http"
	"trabalho/backend/handlers"
)

func StartRoutes() {
	http.HandleFunc("/login", handlers.TokenHandler)
}
