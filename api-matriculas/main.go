package main

import (
	"net/http"
	"trabalho/api-matriculas/routes"

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
