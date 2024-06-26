package endpoint

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/nSakkriou/utils/pkg/agent"
	"github.com/nSakkriou/utils/pkg/logn"
	"github.com/nSakkriou/utils/pkg/util"
)

func GetContentEndpoint(endCommand agent.EndCommand, scriptFolderPath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, GetAllScriptsContent(endCommand, scriptFolderPath))
	}
}

func GetAllScriptsContent(endCommand agent.EndCommand, scriptFolderPath string) string {
	allContent := ""

	for _, script := range endCommand.ScriptsFilesNames {
		allContent += script + "\n"
		allContent += "```\n"
		content, err := getScriptContent(script, scriptFolderPath)
		if err != nil {
			logn.Error("error when reading script file %s", script)
			allContent += err.Error()
		} else {
			logn.Debug("script %s content : %s", script, content)
			allContent += content
		}

		allContent += "\n```"
		allContent += "\n\n"
	}

	return allContent
}

func getScriptContent(scriptFile, scriptFolderPath string) (string, error) {
	scriptFolderPath = util.Prefix(scriptFolderPath, "/")
	scriptFile = strings.TrimPrefix(scriptFile, "/")
	logn.Debug("%s %s", scriptFolderPath, scriptFile)

	fullPath := "." + scriptFolderPath + scriptFile

	contentBytes, err := os.ReadFile(fullPath)
	if err != nil {
		logn.Error("cant read %s %s", fullPath, err)
		return "", nil
	}

	content := string(contentBytes)

	return content, nil
}
