package main

import (
	"fmt"
	"net/http"
	database "trabalho/BD"
	"trabalho/api-matriculas/routes"
)

func main() {
	fmt.Println("Iniciando serviços...")
	database.InitDB()

	defer database.Client.Close()
	routes.SetupRoutes()
	http.ListenAndServe(":8080", nil)

}
