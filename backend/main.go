package main

import (
	"net/http"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/routes"
)

func main() {
	println("Servidor iniciado em localhost:2222/")
	routes.StartRoutes()
	http.ListenAndServe(":2222", nil)
}
