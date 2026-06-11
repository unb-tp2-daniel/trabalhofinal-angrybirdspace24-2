package functions

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"

	// Aqui você importaria os seus arquivos onde estão os códigos dos Handlers
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/functions/admin_functions"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/functions/alunos_functions"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/functions/auth_functions"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/functions/curso_functions"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/functions/departamento_functions"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/functions/materia_functions"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/functions/professor_functions"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/functions/turmas_functions"
)

func init() {

	database.InitDB()

	// Rotas de alunos
	functions.HTTP("ListarAlunos", alunos_functions.ListAlunosHandler)
	functions.HTTP("GetAlunoPorId", alunos_functions.GetAlunoByIdHandler)
	functions.HTTP("CriarAluno", alunos_functions.CreateAlunoHandler)

	// Rotas de turmas
	functions.HTTP("ListarTurmas", turmas_functions.ListTurmasHandler)
	functions.HTTP("CriarTurma", turmas_functions.CreateTurmaHandler)

	// Rotas Materias
	functions.HTTP("CriarMateria", materia_functions.CreateMateriaHandler)
	functions.HTTP("ListarMaterias", materia_functions.ListMateriasHandler)

	// Rotas de Departamento
	functions.HTTP("CriarDepartamento", departamento_functions.CreateDepartamentoHandler)
	functions.HTTP("ListarDepartamentos", departamento_functions.ListDepartamentosHandler)

	// Rotas de Coordenador
	functions.HTTP("CriarCoordenador", admin_functions.CreateCoordenadorHandler)

	// Rotas de curso
	functions.HTTP("CriarCurso", curso_functions.CreateCursoHandler)

	// Rotas de professor
	functions.HTTP("CriarProfessor", professor_functions.CreateProfessorHandler)
	functions.HTTP("ListarProfessores", professor_functions.ListProfessoresHandler)

	functions.HTTP("GetRules", admin_functions.GetRulesHandler)

	//Criando usuário personalizado
	functions.HTTP("CriarUsuario", auth_functions.CreateUser)

	functions.HTTP("MatricularExtraordinaria", alunos_functions.MatriculateAlunoHandler)
	functions.HTTP("Matricular", alunos_functions.NormalMatriculateAlunoHandler)
	functions.HTTP("ProcessarFila", alunos_functions.ProcessarFilaMatriculaHandler)
}
