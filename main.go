package main

import (
	"fmt"
	"github.com/MTthoas/API-Github-GO/functions"
	"github.com/MTthoas/API-Github-GO/model"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"encoding/json"
)

func main() {
	
	fmt.Println("Hello, World!")
	functions.SetLog()

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
		}else{
			log.Info("Récupération des données terminée")
		}

		var repos []model.Repository
		err = json.Unmarshal(body, &repos)
		if err != nil {
			log.Fatal(err)
		}else{
			log.Info("Décodage des données terminé")
		}


		
}
