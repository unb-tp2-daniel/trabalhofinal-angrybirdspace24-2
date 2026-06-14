// functions/alunos_functions/matriculate_aluno.go
package alunos_functions

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	// Importando as nossas pastas isoladas
	"cloud.google.com/go/firestore"
	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/create"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

func MatriculateAlunoHandler(w http.ResponseWriter, r *http.Request) {
	// permite que qualquer origem chegue aqui (ALTERAR DEPOIS; CORS)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// chatice do go
	if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context();

	var matricula models.Matricula
	if err := json.NewDecoder(r.Body).Decode(&matricula); err != nil {
		log.Printf("Erro ao decodificar JSON da matrícula: %v", err)
		http.Error(w, "Formato de dados inválido", http.StatusBadRequest)
		return
	}

	matriculaId := matricula.AlunoId + "_" + matricula.TurmaId
	matriculaRef := database.Client.Collection("matriculas").Doc(matriculaId)
	turmaRef := database.Client.Collection("turmas").Doc(matricula.TurmaId)

	// transação do firebase. Basicamente, impede que outro registro ocorra enquanto esse ta ocorrendo
	// tratando as race conditions

	err := database.Client.RunTransaction(ctx, func(ctx context.Context, t *firestore.Transaction) error {
		matDoc, err := t.Get(matriculaRef)
		if err == nil && matDoc.Exists() {
			return fmt.Errorf("aluno_ja_matriculado")
		}

		turmaDoc, err := t.Get(turmaRef)
		if err != nil {
			return err
		}

		var turma models.Turma
		err = turmaDoc.DataTo(&turma)
		if err != nil {
			return err
		}

		materiaRef := database.Client.Collection("materias").Doc(turma.MateriaId)
		materiaDoc, err := t.Get(materiaRef)
		if (err != nil) {return err}

		var materia models.Materia
		err = materiaDoc.DataTo(&materia)
		if (err != nil) {return err}

		// busca todas as matrículas ativas do aluno
        query := database.Client.Collection("matriculas").Where("alunoId", "==", matricula.AlunoId)
        existingMatriculasIter := t.Documents(query)
        existingMatriculas, err := existingMatriculasIter.GetAll()
        if err != nil {return err}

		for _, docSnap := range existingMatriculas {
            var m models.Matricula
            if err := docSnap.DataTo(&m); err == nil {
                if m.MateriaId == turma.MateriaId {
                    return fmt.Errorf("ja_matriculado_em_outra_turma_da_mesma_materia")
                }
            }
        }

		/* Validação de pré-requisitos */
		alunoRef := database.Client.Collection("alunos").Doc(matricula.AlunoId)
		alunoDoc, err := t.Get(alunoRef)
		if (err != nil) {return err}

		var aluno models.Aluno
		err = alunoDoc.DataTo(&aluno)
		if (err != nil) {return err}

		if aluno.MateriasConcluidas == nil {
            aluno.MateriasConcluidas = make(map[string]string)
        }

		_, concluiu := aluno.MateriasConcluidas[materia.CodigoMateria]
		if (concluiu) {
			return fmt.Errorf("materia_ja_concluida")
		}

		if (len(materia.PreRequisitos) > 0) {
			satisfazUm := false

			for _, req := range materia.PreRequisitos {
				concluiuTodasDoAnd := true

				for _, codigo := range req.Disciplinas {
					_, concluiu := aluno.MateriasConcluidas[codigo]
					
					if (!concluiu) {
						concluiuTodasDoAnd = false
						break // analisa próximo bloco direto
					}
				}

				if (concluiuTodasDoAnd) {
					satisfazUm = true
					break
				}
			}

			if (!satisfazUm) {
				return fmt.Errorf("pre_requisitos_nao_atendidos")
			}			
		}

		/* validação de vagas e update */
		if turma.VagasOcupadas >= turma.VagasTotais {
			return fmt.Errorf("vagas_esgotadas")
		}

		t.Update(turmaRef, []firestore.Update{
			{Path: "vagasOcupadas", Value: turma.VagasOcupadas + 1},
		})

		matricula.Status = true
		matricula.DataSolicitacao = time.Now()

		t.Set(matriculaRef, matricula)

		return nil
	})

	if err != nil {
        switch err.Error() {
        case "aluno_ja_matriculado":
            http.Error(w, "Você já possui uma matrícula ou solicitação ativa nesta turma.", http.StatusConflict)

        case "ja_matriculado_em_outra_turma_da_mesma_materia":
            http.Error(w, "Você já está matriculado em outra turma desta mesma disciplina.", http.StatusConflict)

        case "vagas_esgotadas":
            http.Error(w, "Vagas esgotadas.", http.StatusConflict)

        case "pre_requisitos_nao_atendidos":
            http.Error(w, "Aluno não possui os pré-requisitos necessários.", http.StatusConflict)

        case "materia_ja_concluida":
            http.Error(w, "Aluno já realizou a matéria.", http.StatusConflict)

        default:
            log.Printf("Erro na transação de matrícula: %v", err)
            http.Error(w, "Erro interno ao processar a matrícula", http.StatusInternalServerError)
        }
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Matrícula extraordinária realizada com sucesso!"))
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

	// não há necessidade de verificar se há disponibilidade na turmaref nesse caso

	ctx := r.Context()

	// aplicação de regras de prioridade
	aluno, err := read.GetAlunoById(ctx, database.Client, matricula.AlunoId)
	if err != nil {
		log.Printf("Error ao resgatar aluno no banco de dados %v", err)
		http.Error(w, "Error ao resgatar aluno no banco de dados", http.StatusInternalServerError)
		return
	}

	turmaref, err := read.GetTurmaById(ctx, database.Client, matricula.TurmaId)
	if err != nil {
		log.Printf("Error ao resgatar turmaref no banco de dados %v", err)
		http.Error(w, "Error ao resgatar turmaref no banco de dados", http.StatusInternalServerError)
		return
	}

	materia, err := read.GetMateriaById(ctx, database.Client, turmaref.MateriaId)
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

	err = create.CreateMatricula(ctx, database.Client, matricula)
	if err != nil {
		log.Printf("Erro ao salvar matrícula no banco: %v", err)
		http.Error(w, "Erro interno ao salvar a matrícula", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Matrícula cadastrada com sucesso no banco de dados!"))	
}