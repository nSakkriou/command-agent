package cmd

import (
	"github.com/gorilla/mux"
	"github.com/nSakkriou/command-agent/internal/endpoint"
	"github.com/nSakkriou/command-agent/internal/types"
)

func GetRouter(config *types.Config) *mux.Router {
	router := mux.NewRouter()

	for _, endCommand := range config.EndCommands {
		handler := endpoint.GenerateEndpoint(&endCommand)
		router.HandleFunc(endCommand.EndpointName, handler).Methods(endCommand.Method)
	}

	return router
}
