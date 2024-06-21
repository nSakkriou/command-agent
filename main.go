package main

import (
	"net/http"

	"github.com/nSakkriou/command-agent/cmd"
	"github.com/nSakkriou/command-agent/internal/types"
)

func main() {
	// 1. Charger la config
	// Soit on part du principe quelle est a coté
	// Soit on demande son path en input
	config := types.Config{}

	// 2. Générer le router en étant précautionneur de la validité des infos
	// - Pas de nom en doublons
	// - Pas de fichier qui n'existe pas
	// - Au moins scripts_files_names ou script_file_name de rempli
	// Il faudrait créer une étape de verfication des données
	router := cmd.GetRouter(&config)

	// Démarrage du serveur web avec notre port et notre router
	http.Handle("/", router)
	http.ListenAndServe(":"+config.Port, router)
}
