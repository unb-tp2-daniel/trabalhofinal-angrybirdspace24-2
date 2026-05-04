package routes

import (
	"net/http"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/api-matriculas/handlers"
)

func TeacherRoutes() {
	http.HandleFunc("/teacher/subject/show", handlers.TestHandler)
}
