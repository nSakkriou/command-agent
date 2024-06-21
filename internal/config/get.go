package config

import "github.com/nSakkriou/command-agent/internal/types"

func GetConfig() *types.Config {
	if isLoad {
		return globalConf
	}

	panic("impossible de récupérer la config")
}
