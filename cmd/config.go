package cmd

import (
	"github.com/nSakkriou/command-agent/internal/endpoint"
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
		// Check if endpoint are not allowed (MANDATORY_ENDPOINT)

		if checkMandatoryEndpoint(agentFileConfif) {
			return agentFileConfif
		}
	}

	panic("AgentFile not valid")
}

func checkMandatoryEndpoint(agentFile *agent.AgentFile) bool {
	for _, endCommand := range agentFile.EndCommands {
		for _, notAllowedName := range endpoint.MANDATORY_ENDPOINT {
			if endCommand.EndpointName == notAllowedName {
				logn.Error("endpoint %s is not valid, it part of mandatory endpoint name (%s)", endCommand.EndpointName, endpoint.MANDATORY_ENDPOINT)
				return false
			}
		}
	}

	return true
}
