package materia_functions

import (
	"encoding/json"
	"log"
	"net/http"

	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

// ListTurmasHandler atende a requisição da internet (apenas recebe e devolve)
func ListMateriasHandler(w http.ResponseWriter, r *http.Request) {
	//Tornando o acesso visível para o front
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// verifica se é GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido. Use GET.", http.StatusMethodNotAllowed)
		return
	}

	//chama a função do banco de dados
	materias, err := read.GetAllMaterias(database.Ctx, database.Client)

	if err != nil {
		log.Printf("Erro ao buscar materias no banco: %v", err)
		http.Error(w, "Erro interno ao ler materias", http.StatusInternalServerError)
		return
	}

	//json pro front
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(materias); err != nil {
		log.Printf("Erro ao retornar JSON: %v", err)
		http.Error(w, "Erro na resposta", http.StatusInternalServerError)
	}
}

func ListMateriasObrigatoriasCursoHandler(w http.ResponseWriter, r *http.Request) {
	//Tornando o acesso visível para o front
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// verifica se é GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido. Use GET.", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	curso, err := read.GetCursoById(database.Ctx, database.Client, id)

	var materias []*models.Materia

	for id_materia, _ := range curso.Obrigatorias {	
		materia,err := read.GetMateriaById(database.Ctx, database.Client, id_materia)
		materias = append(materias, materia)
		if err != nil {
			log.Printf("Erro ao buscar materias no banco: %v", err)
			http.Error(w, "Erro interno ao ler materias", http.StatusInternalServerError)
			return
		}
	}

	if err != nil {
		log.Printf("Erro ao buscar materias no banco: %v", err)
		http.Error(w, "Erro interno ao ler materias", http.StatusInternalServerError)
		return
	}

	//json pro front
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(materias); err != nil {
		log.Printf("Erro ao retornar JSON: %v", err)
		http.Error(w, "Erro na resposta", http.StatusInternalServerError)
	}
}

func ListMateriasOptativasCursoHandler(w http.ResponseWriter, r *http.Request) {
	//Tornando o acesso visível para o front
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// verifica se é GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido. Use GET.", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	curso, err := read.GetCursoById(database.Ctx, database.Client, id)

	var materias []*models.Materia

	for id_materia, _ := range curso.Optativas {	
		materia,err := read.GetMateriaById(database.Ctx, database.Client, id_materia)
		materias = append(materias, materia)
		if err != nil {
			log.Printf("Erro ao buscar materias no banco: %v", err)
			http.Error(w, "Erro interno ao ler materias", http.StatusInternalServerError)
			return
		}
	}

	if err != nil {
		log.Printf("Erro ao buscar materias no banco: %v", err)
		http.Error(w, "Erro interno ao ler materias", http.StatusInternalServerError)
		return
	}

	//json pro front
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(materias); err != nil {
		log.Printf("Erro ao retornar JSON: %v", err)
		http.Error(w, "Erro na resposta", http.StatusInternalServerError)
	}
}