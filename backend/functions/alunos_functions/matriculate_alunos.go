// functions/alunos_functions/matriculate_aluno.go
package alunos_functions

import (
	"encoding/json"
	"log"
	"net/http"

	// Importando as nossas pastas isoladas
	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/create"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

// MatriculateAlunoHandler lida exclusivamente com a requisição da internet
func MatriculateAlunoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}
	var matricula models.Matricula
	if err := json.NewDecoder(r.Body).Decode(&matricula); err != nil {
		log.Printf("Erro ao decodificar JSON da matrícula: %v", err)
		http.Error(w, "Formato de dados inválido", http.StatusBadRequest)
		return
	}

	if disponivel, err := read.TurmaIsAvailable(database.Ctx, database.Client, matricula.TurmaId); err != nil {
		http.Error(w, "Erro ao verificar disponibilidade da turma", http.StatusInternalServerError)
		return
	} else if !disponivel {
		http.Error(w, "Turma não disponível para matrícula", http.StatusConflict)
		return
	}

	err := create.CreateMatricula(database.Ctx, database.Client, matricula)
	if err != nil {
		log.Printf("Erro ao salvar matrícula no banco: %v", err)
		http.Error(w, "Erro interno ao salvar a matrícula", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Matrícula cadastrada com sucesso no banco de dados!"))
}

// lida com a matricula normal e rematricula do aluno
func NormalMatriculateAlunoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}
	var matricula models.Matricula
	if err := json.NewDecoder(r.Body).Decode(&matricula); err != nil {
		log.Printf("Erro ao decodificar JSON da matrícula: %v", err)
		http.Error(w, "Formato de dados inválido", http.StatusBadRequest)
		return
	}

	// não há necessidade de verificar se há disponibilidade na turma nesse caso

	ctx := r.Context()

	// aplicação de regras de prioridade
	aluno, err := read.GetAlunoById(ctx, database.Client, matricula.AlunoId)
	if err != nil {
		log.Printf("Error ao resgatar aluno no banco de dados %v", err)
		http.Error(w, "Error ao resgatar aluno no banco de dados", http.StatusInternalServerError)
		return
	}

	turma, err := read.GetTurmaById(ctx, database.Client, matricula.TurmaId)
	if err != nil {
		log.Printf("Error ao resgatar turma no banco de dados %v", err)
		http.Error(w, "Error ao resgatar turma no banco de dados", http.StatusInternalServerError)
		return
	}

	materia, err := read.GetMateriaById(ctx, database.Client, turma.MateriaId)
	if err != nil {
		log.Printf("Error ao resgatar materia no banco de dados %v", err)
		http.Error(w, "Error ao resgatar materia no banco de dados", http.StatusInternalServerError)
		return
	}

	curso, err := read.GetCursoById(ctx, database.Client, aluno.CursoId)
	if err != nil {
		log.Printf("Error ao resgatar curso no banco de dados %v", err)
		http.Error(w, "Error ao resgatar curso no banco de dados", http.StatusInternalServerError)
		return
	}

	matricula.PrioridadeNota = ApplyRules(*aluno, *materia, *curso)

	err = create.CreateMatricula(database.Ctx, database.Client, matricula)
	if err != nil {
		log.Printf("Erro ao salvar matrícula no banco: %v", err)
		http.Error(w, "Erro interno ao salvar a matrícula", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Matrícula cadastrada com sucesso no banco de dados!"))	
}