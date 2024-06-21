package types

type Config struct {
	Port              string
	ScriptsFolderPath string
	EndCommands       []*EndCommand
}
