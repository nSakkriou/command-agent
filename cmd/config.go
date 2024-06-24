package cmd

import (
	"github.com/nSakkriou/utils/pkg/agent"
	"github.com/nSakkriou/utils/pkg/logn"
)

// Load AgentFile and check it
func Config() *agent.AgentFile {
	agentFileConfif, err := agent.Load(agent.AgentFileName)
	if err != nil {
		logn.Error("error on AgentFile loading %s", err)
		panic("cant load and read AgentFile")
	}

	if agent.ValidConfig(agentFileConfif) {
		return agentFileConfif
	}

	panic("AgentFile not valid")
}
