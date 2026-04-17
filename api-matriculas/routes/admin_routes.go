package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func AdminRoutes() {
	http.HandleFunc("/admin", handlers.TesteHandler)

	http.HandleFunc("admin/rules", handlers.TesteHandler)
	http.HandleFunc("admin/rules/add", handlers.TesteHandler)
	http.HandleFunc("admin/rules/delete", handlers.TesteHandler)
	http.HandleFunc("admin/rules/update", handlers.TesteHandler)
}
