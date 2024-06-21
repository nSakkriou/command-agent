package types

type EndCommand struct {
	EndpointName      string   `json:"endpoint_name"`
	Method            string   `json:"method"`
	ScriptsFilesNames []string `json:"scripts_files_names"`
}

// Renvoi script1,script2,script3 false si ils n'existe pas
// Sinon "" true
func (endCommand *EndCommand) CheckIfScriptsExist() (string, bool) {
	return "", false
}

func (endCommand *EndCommand) GetScriptsString() string {
	returned := ""

	for _, script := range endCommand.ScriptsFilesNames {
		returned += script
	}

	return returned
}
