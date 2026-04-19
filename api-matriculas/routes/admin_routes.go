package routes

import (
	"api-matriculas/handlers"
	"net/http"
)

func AdminRoutes() {
	http.HandleFunc("/admin", handlers.TestHandler)

	http.HandleFunc("admin/rules", handlers.TestHandler)
	http.HandleFunc("admin/rules/add", handlers.TestHandler)
	http.HandleFunc("admin/rules/delete", handlers.TestHandler)
	http.HandleFunc("admin/rules/update", handlers.TestHandler)
}
