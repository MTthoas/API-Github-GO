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
	"github.com/gofiber/fiber/v2"
)

func main() {

	functions.SetLog()
	log.Info("Démarrage du programme")

	app := fiber.New()

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

	csvFileName := "./repos/repositories.csv"

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

	functions.ManageCSV(csvFileName, repos)

	log.Info("Fichier CSV créé")

	functions.FetchRepos(user, repos)

	log.Info("Récupération des repos terminée")

	functions.ArchiveRepositories("./repos", "./archives/"+user+".zip")

	app.Get("/download/:username", func(c *fiber.Ctx) error {
		username := c.Params("username")
		filePath := fmt.Sprintf("./archives/%s.zip", username)
	
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return c.Status(404).SendString("Archive not found")
		}
	
		return c.SendFile(filePath)
	})
	


	log.Fatal(app.Listen(":3000"))

}


