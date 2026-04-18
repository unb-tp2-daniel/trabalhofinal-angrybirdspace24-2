package main

import (
	"api-matriculas/routes"
	"net/http"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		println("Erro ao carregar arquivo .env")
		println(err)
		return
	}

	println("Servidor iniciado em localhost:8080/")
	routes.SetupRoutes()
	http.ListenAndServe(":8080", nil)
}
