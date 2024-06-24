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

	// 1. Charger la config (/cmd/config)
	// On va lire un fichier EndCommandFile qui sera sous la forme d'un json
	// On va ensuite valider la config en faisant gaffe à :
	// - tous les champs sont rempli
	// - le dossier de scripts existe
	// - pas de doublons dans le nom des endpoints
	// - les fichiers des endpoints existent
	// Si tout ça est bon, on récupere la config
	logn.Info("Initialisation de la configuration")
	globalConf := cmd.Config()

	// 2. Générer le router et tous les endpoints défini dans le config (/cmd/router)
	logn.Info("Création du router")
	router := cmd.GetRouter(globalConf)

	// 3. Démarrage du serveur web avec notre port et notre router
	http.Handle("/", router)

	logn.Info("Démarrage du serveur ... port %d", globalConf.Port)
	err := http.ListenAndServe(":"+fmt.Sprint(globalConf.Port), router)
	if err != nil {
		logn.Error("impossible de demarrer le serveur %s", err)
		os.Exit(0)
	}

}

func handlePanic() {
	if r := recover(); r != nil {
		logn.Error("Impossible de continuer : %v", r)
	}
}
