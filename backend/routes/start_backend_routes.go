package routes

import (
	"net/http"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/handlers"
	firebase "firebase.google.com/go/v4"
)

func StartRoutes(app *firebase.App) {
	http.HandleFunc("/login", handlers.TokenHandler)

	/* precisam de proteção */
	StudentRoutes(app)
	AdminRoutes(app)
	/* precisam de proteção */

	GeneralRoutes()
	TeacherRoutes()
	DirectorRoutes()
	DatabaseRoutes()
}
