package functions

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	database "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/BD"

	// Aqui você importaria os seus arquivos onde estão os códigos dos Handlers
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/functions/alunos_functions"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/functions/materia_functions"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/functions/turmas_functions"

	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/functions/admin_functions"
	"github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/functions/departamento_functions"
)

func init() {
	// 1. Iniciamos o banco de dados (que você configurou perfeitamente no passo anterior)
	database.InitDB()

	// 2. O ROTEAMENTO SERVERLESS:
	// Aqui nós dizemos ao Google Cloud: "Pegue estas funções em Go e transforme em URLs independentes"

	// Rota de Login (que estava solta no seu código original)
	//functions.HTTP("Login", TokenHandler) // O Google criará uma URL terminada em /Login

	// Rotas de Admin (o que antes estava dentro de AdminRoutes)
	//functions.HTTP("CriarTurmaAdmin", CriarTurmaHandler)
	//functions.HTTP("DeletarTurmaAdmin", DeletarTurmaHandler)

	// Rotas de Student (o que antes estava dentro de StudentRoutes)
	//functions.HTTP("MatricularAluno", MatricularHandler)
	//functions.HTTP("RetirarMatricula", RetirarMatriculaHandler)

	// Rotas de Database/Listagem
	functions.HTTP("ListarTurmas", turmas_functions.ListTurmasHandler)
	functions.HTTP("CriarTurma", turmas_functions.CreateTurmaHandler)
	functions.HTTP("CriarMateria", materia_functions.CreateMateriaHandler)
	functions.HTTP("CriarDepartamento", departamento_functions.CreateDepartamentoHandler)

	functions.HTTP("GetRules", admin_functions.GetRulesHandler)

	functions.HTTP("MatricularExtraordinaria", alunos_functions.MatriculateAlunoHandler)
	functions.HTTP("Matricular", alunos_functions.NormalMatriculateAlunoHandler)
}