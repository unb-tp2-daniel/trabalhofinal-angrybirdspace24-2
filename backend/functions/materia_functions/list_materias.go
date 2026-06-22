package materia_functions

import (
	"encoding/json"
	"log"
	"net/http"

	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	cursoDB "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read/curso"
	materiaDB "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read/materia"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func ListMateriasHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	// verifica se é GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido. Use GET.", http.StatusMethodNotAllowed)
		return
	}

	//chama a função do banco de dados
	materias, err := materiaDB.GetAllMaterias(database.Ctx, database.Client)

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

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido. Use GET.", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	curso, err := cursoDB.GetCursoById(database.Ctx, database.Client, id)

	var materias []*models.Materia

	for id_materia, _ := range curso.Obrigatorias {
		materia, err := materiaDB.GetMateriaById(database.Ctx, database.Client, id_materia)
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
	curso, err := cursoDB.GetCursoById(database.Ctx, database.Client, id)

	var materias []*models.Materia

	for id_materia, _ := range curso.Optativas {
		materia, err := materiaDB.GetMateriaById(database.Ctx, database.Client, id_materia)
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

// DataRetrievalMateriasHandler retorna todas as matérias com todos os atributos.
// Rota exclusiva para administração/data retrieval, sem interferir no front-end.
func DataRetrievalMateriasHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Verifica se é GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido. Use GET.", http.StatusMethodNotAllowed)
		return
	}

	// Chama a função do banco de dados (que já resgata todos os atributos do model)
	materias, err := materiaDB.GetAllMaterias(database.Ctx, database.Client)

	if err != nil {
		log.Printf("Erro de Data Retrieval ao buscar materias: %v", err)
		http.Error(w, "Erro interno ao ler materias", http.StatusInternalServerError)
		return
	}

	// Retorna o JSON completo
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(materias); err != nil {
		log.Printf("Erro ao retornar JSON no Data Retrieval: %v", err)
		http.Error(w, "Erro na resposta", http.StatusInternalServerError)
	}
}
