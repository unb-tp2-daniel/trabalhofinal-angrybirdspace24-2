package routes

import (
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/middleware"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/handlers"
)

func StudentRoutes(app *firebase.App) {
	// nessas rotas vão ser usadas middleware pelo firebase
	http.HandleFunc("/student", handlers.PingHandler)
	http.HandleFunc("/student/matriculate", middleware.AuthMiddleware(app, handlers.PingHandler))
	http.HandleFunc("/student/position", handlers.PingHandler)
	http.HandleFunc("/student/subjects", handlers.PingHandler)
	http.HandleFunc("/student/schedule", handlers.PingHandler)
	//http.HandleFunc("/student/matriculate/extraordinary", handlers.TesteHandler)
	//http.HandleFunc("/student/matriculate/regular", handlers.TesteHandler)
}
