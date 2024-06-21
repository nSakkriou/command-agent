package config

import (
	"fmt"
	"path/filepath"

	"github.com/nSakkriou/command-agent/internal/types"
	"github.com/nSakkriou/command-agent/internal/util"
)

// Checker le fichier de config
// Tous les champs rempli
// Au moins 1 script dans la liste
// Pas de doublons dans les endpoints
// checker si les fichiers existes
func ValidConfig(config *types.Config) bool {
	if config.Port == 0 || config.Port == -1 {
		fmt.Println("le n° de port n'est pas correct")
		return false
	}

	if util.IsEmpty(config.ScriptsFolderPath) {
		fmt.Println("le champ scripts_folder_path est vide")
		return false
	}

	// Check si le dossier existe
	if util.FileExist(config.ScriptsFolderPath) {
		fmt.Println(config.ScriptsFolderPath + " does not exist")
		return false
	}

	endpointNames := []string{}
	for _, endCommand := range config.EndCommands {
		// Checker les doublons
		if util.ArrayContains(endpointNames, endCommand.EndpointName) {
			fmt.Println(endCommand.EndpointName + " est en double")
			return false
		}

		endpointNames = append(endpointNames, endCommand.EndpointName)

		// Checker si les fichiers existe
		for _, scriptName := range endCommand.ScriptsFilesNames {
			path := filepath.Join(config.ScriptsFolderPath, scriptName)
			if util.FileExist(path) {
				fmt.Println(path + "does not exist")
				return false
			}
		}
	}

	return true
}
