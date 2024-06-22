package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nSakkriou/command-agent/cmd"
	"github.com/nSakkriou/command-agent/internal/util"
	"github.com/theritikchoure/logx"
)

func main() {
	logx.ColoringEnabled = true
	defer handlePanic()

	// 1. Charger la config (/cmd/config)
	// On va lire un fichier EndCommandFile qui sera sous la forme d'un json
	// On va ensuite valider la config en faisant gaffe à :
	// - tous les champs sont rempli
	// - le dossier de scripts existe
	// - pas de doublons dans le nom des endpoints
	// - les fichiers des endpoints existent
	// Si tout ça est bon, on récupere la config
	util.Info("Initialisation de la configuration", "")
	globalConf := cmd.Config()

	// 2. Générer le router et tous les endpoints défini dans le config (/cmd/router)
	util.Info("Création du router", "")
	router := cmd.GetRouter(globalConf)

	// 3. Démarrage du serveur web avec notre port et notre router
	http.Handle("/", router)

	util.Info("Démarrage du serveur ... port %d", globalConf.Port)
	err := http.ListenAndServe(":"+fmt.Sprint(globalConf.Port), router)
	if err != nil {
		util.Error("impossible de demarrer le serveur %s", err)
		os.Exit(0)
	}

}

func handlePanic() {
	if r := recover(); r != nil {
		util.Error("Impossible de continuer : %v", r)
	}
}
