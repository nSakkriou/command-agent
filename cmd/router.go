package cmd

import (
	"strings"

	"github.com/gorilla/mux"
	"github.com/nSakkriou/command-agent/internal/endpoint"
	"github.com/nSakkriou/command-agent/internal/types"
	"github.com/nSakkriou/command-agent/internal/util"
	"github.com/theritikchoure/logx"
)

func GetRouter(config *types.Config) *mux.Router {
	router := mux.NewRouter()

	for _, endCommand := range config.EndCommands {
		util.Debug("description end command : %s", logx.FGBLUE, endCommand)

		handler := endpoint.GenerateEndpoint(endCommand, config.ScriptsFolderPath)
		router.HandleFunc(util.Prefix(endCommand.EndpointName, "/"), handler).Methods(endCommand.Method)
	}

	util.Verbose("Descriptif des routes")

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, err1 := route.GetPathTemplate()
		met, err2 := route.GetMethods()

		if err1 != nil && err2 != nil {
			util.Error("erreurs : %s - %s", err1, err2)
			panic("l'un des endpoints n'est pas correct")
		} else {
			util.Verbose("%s %s", tpl, strings.Join(met, ","))
		}
		return nil
	})

	return router
}
