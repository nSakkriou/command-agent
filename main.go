package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nSakkriou/command-agent/cmd"
	"github.com/nSakkriou/utils/pkg/logn"
)

func main() {
	defer handlePanic()

	// 1. load AgentFile (/cmd/config)
	// AgentFile is Json file
	// We gonna check some things :
	// - no empty field
	// - scripts folder exists
	// - no duplicate endpoint
	// - endpoint script existe
	// if everythings is ok, we return config stuct
	logn.Info("load and check AgentFile")
	globalConf := cmd.Config()

	// 2. Generate router with basic and custom endpoint (/cmd/router)
	logn.Info("generate custom router")
	router := cmd.GetRouter(globalConf)

	// 3. start web server with custom router and port
	http.Handle("/", router)

	logn.Info("server start ... port %d", globalConf.Port)
	err := http.ListenAndServe(":"+fmt.Sprint(globalConf.Port), router)
	if err != nil {
		logn.Error("cant start server %s", err)
		os.Exit(0)
	}

}

func handlePanic() {
	if r := recover(); r != nil {
		logn.Error("Impossible de continuer : %v", r)
	}
}
