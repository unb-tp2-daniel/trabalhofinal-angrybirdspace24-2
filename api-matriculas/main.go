package main

import (
	"fmt"
	"net/http"
	database "trabalho/BD"
	"trabalho/api-matriculas/routes"
  "github.com/joho/godotenv"
)

func main() {
	fmt.Println("Iniciando serviços...")
	database.InitDB()
	database.SeedBaseData()

	defer database.Client.Close()
	routes.SetupRoutes()
	http.ListenAndServe(":8080", nil)

}
