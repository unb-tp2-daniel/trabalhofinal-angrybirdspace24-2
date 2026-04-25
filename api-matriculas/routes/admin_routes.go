package routes

import (
	"net/http"
	"trabalho/api-matriculas/handlers"
)

func AdminRoutes() {
	http.HandleFunc("/admin", handlers.TestHandler)

	http.HandleFunc("admin/rules", handlers.TestHandler)
	http.HandleFunc("admin/rules/add", handlers.TestHandler)
	http.HandleFunc("admin/rules/delete", handlers.TestHandler)
	http.HandleFunc("admin/rules/update", handlers.TestHandler)
}
