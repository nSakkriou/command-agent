package cmd

import (
	"github.com/nSakkriou/command-agent/internal/config"
	"github.com/nSakkriou/command-agent/internal/types"
)

func Config() *types.Config {
	config.Load()

	globalConf := config.GetConfig()

	if config.ValidConfig(globalConf) {
		return globalConf
	}

	panic("la configuration n'est pas valide")
}
