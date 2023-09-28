package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MTthoas/API-Github-GO/functions"
	"github.com/MTthoas/API-Github-GO/model"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"encoding/csv"
)

func main() {

	functions.SetLog()
	log.Info("Démarrage du programme")

	// Charger le fichier .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur lors de la lecture du fichier .env")
	}

	dir := "./repos"

	errRemoving := functions.RemoveContents(dir)

	if err != nil {
		log.Fatal(errRemoving)
	}

	user := os.Getenv("GITHUB_USER")
	token := os.Getenv("GITHUB_TOKEN")

	// Spécification à 100 pages piles
	url := fmt.Sprintf("https://api.github.com/users/%s/repos?per_page=100", user)
	body, err := functions.GetMethod(url, token)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Récupération des données terminée")
	}

	var repos []model.Repository
	err = json.Unmarshal(body, &repos)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Décodage des données terminé")
	}

	repos = functions.SortByDate(repos)

	csvFileName := "repositories.csv"

	// Supprimer le fichier CSV s'il existe
	if _, err := os.Stat(csvFileName); !os.IsNotExist(err) {
		os.Remove(csvFileName)
		log.Info("Fichier CSV existant supprimé")
	}

	// Créer un nouveau fichier CSV
	csvFile, err := os.Create(csvFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	functions.ExcelWrite(repos, writer)

	log.Info("Fichier CSV créé")

	functions.FetchRepos(user, repos)

	log.Info("Récupération des repos terminée")

}
