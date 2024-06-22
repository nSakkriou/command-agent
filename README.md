Outil permettant de générer à partir d'un fichier de config un api web, qui lie un endpoint avec un script shell

Au démarrage :
agent path/config.json

Charge la config
Génére les endpoints
    - Check si le script existe

---

Le config.json doit s'appeller AgentFile | EndCommandFile, ça ne sera que du json dedans, mais au moins comme ça on ne se pose pas de question du nom

TODO :
- Fixer les logs (artefacts)
- Tester sur ma prod
- Créer les commandes dans ma cli