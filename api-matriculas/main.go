package main

import (
	"fmt"
	"net/http"

	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/BD"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/api-matriculas/routes"
	//"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Iniciando serviços...")
	database.InitDB()
	database.SeedBaseData()

	defer database.Client.Close()
	routes.SetupRoutes()
	http.ListenAndServe(":8080", nil)

}
