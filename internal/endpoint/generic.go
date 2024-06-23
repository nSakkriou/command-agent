package endpoint

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"

	"github.com/nSakkriou/command-agent/internal/types"
	"github.com/nSakkriou/command-agent/internal/util"
)

// Générer les fonctions des endpoint à partir d'un EndCommand
func GenerateEndpoint(endCommand types.EndCommand, scriptFolderPath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w = setHeader(w)

		response := types.JSONResponse{
			Endpoint: endCommand.EndpointName,
			Outputs:  []types.FileOutput{},
		}

		util.Info("Execution du endpoint : %s", response.Endpoint)

		// Execution des scripts
		for _, script := range endCommand.ScriptsFilesNames {
			output := types.FileOutput{
				Filename: script,
				Output:   "null",
				Success:  true,
			}

			outputCommand, err := execCommand(script, scriptFolderPath)
			if err != nil {
				outputCommand = outputCommand + " " + err.Error()
				output.Success = false
			}

			output.Output = outputCommand

			util.Verbose("Résulat du script : %s", script)
			util.Verbose("%s", output)

			response.Outputs = append(response.Outputs, output)
		}

		output(w, response)
	}
}

func setHeader(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")

	return w
}

func output(w http.ResponseWriter, response interface{}) {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func execCommand(scriptFile string, scriptFolderPath string) (string, error) {
	scriptFolderPath = util.Prefix(scriptFolderPath, "/")
	scriptFile = strings.TrimPrefix(scriptFile, "/")

	util.Debug("%s %s", scriptFolderPath, scriptFile)

	fullPath := scriptFolderPath + scriptFile

	fullCommand := "." + fullPath

	util.Verbose("command %s", fullCommand)

	cmd := exec.Command("sh", "-c", fullCommand)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}

	return string(output), nil
}
