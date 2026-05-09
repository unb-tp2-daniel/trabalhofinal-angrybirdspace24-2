package routes

import (
	"net/http"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/handlers"
)

func TeacherRoutes() {
	http.HandleFunc("/teacher/subject/show", handlers.PingHandler)
}
