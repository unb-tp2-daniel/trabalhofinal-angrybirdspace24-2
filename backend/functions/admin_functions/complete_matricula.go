package admin_functions

import (
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

type CandidatoMatricula struct {
	Matricula models.Matricula
	DocRef    *firestore.DocumentRef
}

// depois do periodo de matricula/rematricula, essa função é rodada para ver quem passou ou não

/*
	ESSA FUNÇÃO VAI SE TORNAR UM GSCHEDULE
	isso significa que ela vai rodar sozinha agendada em algum momento. Com essa configuração,
	é possível aumentar o timeout para 2 horas. Como, por ora, ela está em HTTP, então o tempo padrão,
	estourará com certeza com 22k de turmas
*/

func CompleteMatriculaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Admin-Token")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido. Use POST para processar as matrículas.", http.StatusMethodNotAllowed)
		return
	}

	adminToken := r.Header.Get("X-Admin-Token")
	if adminToken == "" || adminToken != TokenAdminSecreto {
		log.Printf("Tentativa de reset não autorizada vinda do IP: %s", r.RemoteAddr)
		http.Error(w, "Não autorizado. Token administrativo inválido.", http.StatusUnauthorized)
		return
	}

	ctx := r.Context()
	client := database.Client

	turmasDoc, err := client.Collection("turmas").Where("ativo", "==", true).Documents(ctx).GetAll()
	//ids := []interface{}{"ADM01_0004_05", "DI01_0014_06"}
	//turmasDoc, err := client.Collection("turmas").Where("codigoTurma", "in", ids).Documents(ctx).GetAll()
	if err != nil {
		log.Printf("Erro crítico ao buscar turmas ativas: %v", err)
		http.Error(w, "Erro crítico ao buscar turmas ativas: ", http.StatusInternalServerError)
		return
	}

	for _, turmaDoc := range turmasDoc {
		var turma models.Turma
		if err := turmaDoc.DataTo(&turma); err != nil {
			log.Printf("Erro ao decodificar turma %s, pulando...", turmaDoc.Ref.ID)
			continue
		}

		turmaId := turmaDoc.Ref.ID
		batch := client.Batch()

		query := client.Collection("matriculas").
			Where("turmaId", "==", turmaId).
			OrderBy("prioridadenota", firestore.Desc).
			OrderBy("dataSolicitacao", firestore.Asc)

		matDocs, err := query.Documents(ctx).GetAll()
		if err != nil {
			log.Printf("Erro ao buscar matriculas da turma %s: %v", turmaId, err)
			continue
		}

		if len(matDocs) == 0 { // salva o algoritmo de explodir o tempo
			continue
		}

		var fila []CandidatoMatricula

		for _, doc := range matDocs {
			var m models.Matricula
			if err := doc.DataTo(&m); err != nil {
				continue
			}

			fila = append(fila, CandidatoMatricula{
				Matricula: m,
				DocRef:    doc.Ref,
			})
		}

		ganhou := make(map[int]bool)
		vagasExclusivasUsadas := make(map[string]int64)

		// alocação das vagas exclusivas
		for i, candidato := range fila {
			limite, temVagaEx := turma.VagasExclusivas[candidato.Matricula.CursoId]

			if temVagaEx {
				atual := vagasExclusivasUsadas[candidato.Matricula.CursoId]
				if atual < limite {
					ganhou[i] = true
					vagasExclusivasUsadas[candidato.Matricula.CursoId]++
				}
			}
		}

		var totalExPreenchidas int64 = 0
		for _, qtd := range vagasExclusivasUsadas {
			totalExPreenchidas += qtd
		}

		// alocação das vagas universais restantes
		vagasRestantes := turma.VagasTotais - totalExPreenchidas

		for i := range fila {
			if ganhou[i] {
				continue
			}

			if vagasRestantes > 0 {
				ganhou[i] = true
				vagasRestantes--
			}
		}

		var totalOcupadas int64 = 0

		for i, candidato := range fila {
			statusFinal := ganhou[i]

			batch.Update(candidato.DocRef, []firestore.Update{
				{Path: "status", Value: statusFinal},
			})

			if statusFinal {
				totalOcupadas++
			}
		}

		turmaRef := client.Collection("turmas").Doc(turmaId)
		batch.Update(turmaRef, []firestore.Update{
			{Path: "vagasOcupadas", Value: totalOcupadas},
		})

		// commita as alterações da turma inteira de uma só vez
		_, err = batch.Commit(ctx)
		if err != nil {
			log.Printf("Erro crítico ao commitar batch da turma %s: %v", turmaId, err)
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"mensagem": "Período de matrículas encerrado e vagas distribuídas com sucesso!"}`))
}