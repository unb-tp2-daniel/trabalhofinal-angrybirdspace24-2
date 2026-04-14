package routes

import ("net/http"; "api-matriculas/handlers")
func SetupRoutes() 
{ 
	http.HandleFunc("/users", handlers.UsersHandler) 
}