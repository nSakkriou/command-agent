package cmd

import (
	"strings"

	"github.com/gorilla/mux"
	"github.com/nSakkriou/command-agent/internal/endpoint"
	"github.com/nSakkriou/utils/pkg/agent"
	"github.com/nSakkriou/utils/pkg/logn"
	"github.com/nSakkriou/utils/pkg/util"
)

func GetRouter(config *agent.AgentFile) *mux.Router {
	router := mux.NewRouter()

	for _, endCommand := range config.EndCommands {
		logn.Debug("description end command : %+v", endCommand)

		handler := endpoint.GenerateEndpoint(endCommand, config.ScriptsFolderPath)
		router.HandleFunc(util.Prefix(endCommand.EndpointName, "/"), handler).Methods(endCommand.Method)
	}

	// Route de base
	router.HandleFunc("/available", endpoint.AvailableEndpoint(config.EndCommands)).Methods("GET")
	router.HandleFunc("/health", endpoint.HealthEndpoint).Methods("GET")

	logn.Verbose("Descriptif des routes")

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, err1 := route.GetPathTemplate()
		met, err2 := route.GetMethods()

		if err1 != nil && err2 != nil {
			logn.Error("erreurs : %s - %s", err1, err2)
			panic("l'un des endpoints n'est pas correct")
		} else {
			logn.Verbose("%s %s", tpl, strings.Join(met, ","))
		}
		return nil
	})

	return router
}
