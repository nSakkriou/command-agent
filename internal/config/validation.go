package config

import (
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
		util.Error("le n° de port n'est pas correct")
		return false
	}

	if util.IsEmpty(config.ScriptsFolderPath) {
		util.Error("le champ scripts_folder_path est vide")
		return false
	}

	// Check si le dossier existe
	if util.FileExist(config.ScriptsFolderPath) {
		util.Error("le dossier " + config.ScriptsFolderPath + " n'existe pas")
		return false
	}

	endpointNames := []string{}
	for _, endCommand := range config.EndCommands {
		// Checker les doublons
		if util.ArrayContains(endpointNames, endCommand.EndpointName) {
			util.Error("le endpoint " + endCommand.EndpointName + " ne peut pas exister en double")
			return false
		}

		endpointNames = append(endpointNames, endCommand.EndpointName)

		// Checker si les fichiers existe
		for _, scriptName := range endCommand.ScriptsFilesNames {
			path := filepath.Join(config.ScriptsFolderPath, scriptName)
			if util.FileExist(path) {
				util.Error("le script " + path + " n'existe pas")
				return false
			}
		}
	}

	return true
}
