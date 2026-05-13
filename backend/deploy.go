package functions

import (
	"fmt"
	"os"
	"os/exec"
)

var nomesfuncoes = []string{
	"ListarTurmas",
}

func deployFunction(name string) error {

	args := []string{
		"functions",
		"deploy",
		name,
		"--gen2",
		"--runtime=go126",
		"--region=southamerica-east1",
		"--source=.",
		"--entry-point=" + name,
		"--trigger-http",
		"--allow-unauthenticated",
		"--set-build-env-vars=GOFLAGS=-buildvcs=false",
		"--project=matriculas242",
	}

	cmd := exec.Command("gcloud", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func main() {

	for _, fn := range nomesfuncoes {

		fmt.Println("Deployando:", fn)

		err := deployFunction(fn)

		if err != nil {
			fmt.Println("Erro:", err)
			return
		}
	}

	fmt.Println("Todas deployadas.")
}
