package endpoint

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"

	"github.com/nSakkriou/command-agent/internal/types"
	"github.com/nSakkriou/utils/pkg/agent"
	"github.com/nSakkriou/utils/pkg/logn"
	"github.com/nSakkriou/utils/pkg/util"
)

// Generate custom endpoint with agent.EndCommand information
func GenerateEndpoint(endCommand agent.EndCommand, scriptFolderPath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w = setHeader(w)

		response := types.JSONResponse{
			Endpoint: endCommand.EndpointName,
			Outputs:  []types.FileOutput{},
		}

		logn.Info("Endpoint %s execution", response.Endpoint)

		// scripts exec
		for _, script := range endCommand.ScriptsFilesNames {
			output := types.FileOutput{
				Filename: script,
				Output:   "",
				Success:  true,
			}

			outputCommand, err := execCommand(script, scriptFolderPath)
			if err != nil {
				outputCommand = outputCommand + "\n" + err.Error()
				output.Success = false
			}

			output.Output = outputCommand

			logn.Verbose("Script result : %s", script)
			logn.Verbose("%s", output)

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

func execCommand(scriptFile, scriptFolderPath string) (string, error) {
	scriptFolderPath = util.Prefix(scriptFolderPath, "/")
	scriptFile = strings.TrimPrefix(scriptFile, "/")
	logn.Debug("%s %s", scriptFolderPath, scriptFile)

	fullPath := scriptFolderPath + scriptFile
	fullCommand := "." + fullPath
	logn.Verbose("command %s", fullCommand)

	cmd := exec.Command("sh", "-c", fullCommand)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}

	logn.Debug("output %s", string(output))
	return string(output), nil
}
