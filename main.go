package main

import (
	"encoding/json"
	"fmt"
	"os"

	"encoding/csv"
	"log"
	"github.com/MTthoas/API-Github-GO/functions"
	"github.com/MTthoas/API-Github-GO/model"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func main() {

	functions.SetLog()
	log.Println("Démarrage du programme")

	// Créez les dossiers si ils n'existent pas
	if err := os.MkdirAll("./repos", os.ModePerm); err != nil {
		log.Fatal("Erreur lors de la création du dossier /repos :", err)
	}
	if err := os.MkdirAll("./archives", os.ModePerm); err != nil {
		log.Fatal("Erreur lors de la création du dossier /archives :", err)
	}

	app := fiber.New()

	log.Println("Démarrage de l'application localhost:3000")

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
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", user)
	body, err := functions.GetMethod(url, token)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Récupération des données terminée")
	}

	var repos []model.Repository
	err = json.Unmarshal(body, &repos)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Décodage des données terminé")
	}

	repos = functions.SortByDate(repos)
	if len(repos) > 100 {
		repos = repos[:100]
	}

	csvFileName := "./repos/repositories.csv"

	// Supprimer le fichier CSV s'il existe
	if _, err := os.Stat(csvFileName); !os.IsNotExist(err) {
		os.Remove(csvFileName)
		log.Println("Fichier CSV existant supprimé")
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

	log.Println("Fichier CSV créé")

	functions.FetchRepos(user, repos)

	log.Println("Récupération des repos terminée")

	functions.ArchiveRepositories("./repos", "./archives/"+user+".zip")

	log.Println("Archivage des repos terminée")

	app.Get("/download/:username", func(c *fiber.Ctx) error {
		username := c.Params("username")
		filePath := fmt.Sprintf("./archives/%s.zip", username)

		log.Println("Téléchargement de l'archive :", filePath)
		log.Println(("Ip du client :"), c.IP())

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return c.Status(404).SendString("Archive not found")
		}

		return c.SendFile(filePath)
	})

	log.Fatal(app.Listen(":3000"))

}
