package main

import (
	"api-matriculas/routes"
	"net/http"
)

func main() {
	routes.SetupRoutes()
	http.ListenAndServe(":8080", nil)
}
