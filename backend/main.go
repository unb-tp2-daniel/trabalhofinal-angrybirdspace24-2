package main

import (
	"net/http"
	"trabalho/backend/routes"
)

func main() {
	println("Servidor iniciado em localhost:2222/")
	routes.StartRoutes()
	http.ListenAndServe(":2222", nil)
}
