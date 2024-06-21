package endpoint

import (
	"fmt"
	"net/http"

	"github.com/nSakkriou/command-agent/internal/types"
)

// Générer les fonctions des endpoint à partir d'un EndCommand
func GenerateEndpoint(endCommand *types.EndCommand) func(w http.ResponseWriter, r *http.Request) {

	scriptsDoestExist, status := endCommand.CheckIfScriptsExist()

	if status {
		return func(w http.ResponseWriter, r *http.Request) {
			// Execution des scripts
			fmt.Fprintf(w, endCommand.GetScriptsString())
		}
	} else {
		panic("impossible de générer le endpoint " + endCommand.EndpointName + ". les scripts " + scriptsDoestExist + " n'existe pas")
	}
}
