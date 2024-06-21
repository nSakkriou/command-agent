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

	// 1. Charger la config
	// Soit on part du principe quelle est a coté
	// Soit on demande son path en input
	util.Info("Initialisation de la configuration", "")
	globalConf := cmd.Config()

	// 2. Générer le router en étant précautionneur de la validité des infos
	// - Pas de nom en doublons
	// - Pas de fichier qui n'existe pas
	// - Au moins scripts_files_names ou script_file_name de rempli
	// Il faudrait créer une étape de verfication des données
	util.Info("Création du router", "")
	router := cmd.GetRouter(globalConf)

	// Démarrage du serveur web avec notre port et notre router
	http.Handle("/", router)

	util.Info("Démarrage du serveur ... port %d", globalConf.Port)
	err := http.ListenAndServe(":"+fmt.Sprint(globalConf.Port), router)
	if err != nil {
		util.Error("impossible de demarrer le serveur %s", err)
		os.Exit(0)
	}

}
