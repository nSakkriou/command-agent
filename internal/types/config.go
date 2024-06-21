package types

type Config struct {
	Port              int          `json:"port"`
	ScriptsFolderPath string       `json:"scripts_folder_path"`
	EndCommands       []EndCommand `json:"endcommands"`
}
