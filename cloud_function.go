package p // O pacote não precisa ser main para o Cloud Functions

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	// Registra o ponto de entrada da Function
	// "ApiMatriculas" é o nome que usaremos na hora do deploy
	functions.HTTP("ApiMatriculas", HandlerNuvem)
}

// HandlerNuvem é a função que receberá o tráfego da nuvem
func HandlerNuvem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Function implantada direto na arquitetura do projeto de matrículas!")
}
