package routes

import (
	"net/http"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/handlers"
	firebase "firebase.google.com/go/v4"
)

func AdminRoutes(app *firebase.App) {
	// nessas rotas vão ser usadas middleware pelo firebase
	http.HandleFunc("/admin", handlers.PingHandler)

	http.HandleFunc("/admin/rules", handlers.PingHandler)
	http.HandleFunc("/admin/rules/add", handlers.PingHandler)
	http.HandleFunc("/admin/rules/delete", handlers.PingHandler)
	http.HandleFunc("/admin/rules/update", handlers.PingHandler)
}
