package cmd

import (
	"github.com/nSakkriou/utils/pkg/agent"
	"github.com/nSakkriou/utils/pkg/logn"
)

func Config() *agent.AgentFile {
	agentFileConfif, err := agent.Load(agent.AgentFileName)
	if err != nil {
		logn.Error("erreur lors du load de l'AgentFile %s", err)
		panic("impossible de lire de charger l'agent file")
	}

	if agent.ValidConfig(agentFileConfif) {
		return agentFileConfif
	}

	panic("la configuration n'est pas valide")
}
