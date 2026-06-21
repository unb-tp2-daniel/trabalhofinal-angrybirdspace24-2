package admin_functions

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	alunoDB "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD/read/aluno"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/models"
)

type CandidatoMatricula struct {
	Matricula models.Matricula
	DocRef    *firestore.DocumentRef
	CursoId   string
}

// depois do periodo de matricula/rematricula, essa função é rodada para ver quem passou ou não
func CompleteMatricula(ctx context.Context, client *firestore.Client) error {
	// pega todas as turmas ativas
	turmasDoc, err := client.Collection("turmas").Where("ativo", "==", true).Documents(ctx).GetAll()

	if err != nil {
		return nil
	}

	for _, turmaDoc := range turmasDoc {
		var turma models.Turma
		err = turmaDoc.DataTo(&turma)
		if err != nil {
			continue
		}

		turmaId := turmaDoc.Ref.ID // pega id do documento em si ja

		batch := client.Batch() // permite que tenha varias operações sem modificar o banco de dados de vdd

		query := client.Collection("matriculas").
			Where("turmaId", "==", turmaId).
			OrderBy("prioridadeNota", firestore.Desc).
			OrderBy("dataSolicitacao", firestore.Asc)

		matDocs, err := query.Documents(ctx).GetAll() // resgata o resultado da query

		if err != nil {
			log.Printf("Erro ao buscar matriculas da turma %s: %v", turmaId, err)
			continue
		}

		// auxiliares
		var fila []CandidatoMatricula

		for _, doc := range matDocs {
			var m models.Matricula
			doc.DataTo(&m)

			/* Talvez usar outra abordagem para isso. Com o tempo, e muitas requisições, pode trazer lentidão.
			A outra abordagem que pensei foi em criar um CursoId no model Matricula, talvez ajude, mas não sei onde colocaar*/
			aluno, err := alunoDB.GetAlunoById(ctx, client, m.AlunoId)
			if err != nil {
				continue
			}

			fila = append(fila, CandidatoMatricula{
				Matricula: m,
				DocRef:    doc.Ref,
				CursoId:   aluno.CursoId,
			})
		}

		// controlar quem ja ganhou vaga e quantas vagas exclusivas foram usadas
		ganhou := make(map[int]bool) // mapeia o indice do aluno na lista original
		vagasExclusivasUsadas := make(map[string]int64)

		// assegurar as vagas exclusivas
		for i, candidato := range fila {
			limite, temVagaEx := turma.VagasExclusivas[candidato.CursoId]

			if temVagaEx {
				atual := vagasExclusivasUsadas[candidato.CursoId]
				if atual < limite {
					ganhou[i] = true
					vagasExclusivasUsadas[candidato.CursoId]++
				}
			}
		}

		// aplicar o restante das vagas
		var totalExPreenchidas int64 = 0
		for _, qtd := range vagasExclusivasUsadas {
			totalExPreenchidas += qtd
		}

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
				{Path: "status", Value: statusFinal}, // aceita a matricula dele
			})

			if statusFinal {
				totalOcupadas++
			}
		}

		// atualiza o metadata
		turmaRef := client.Collection("turmas").Doc(turmaId)
		batch.Update(turmaRef, []firestore.Update{
			{Path: "VagasOcupadas", Value: totalOcupadas},
		})

		// commita as atualizações de uma só vez
		_, err = batch.Commit(ctx)
		if err != nil {
			log.Printf("erro ao fechar a turma %s: %v", turmaId, err)
		}
	}

	return nil
}
