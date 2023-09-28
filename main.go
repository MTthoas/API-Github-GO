package main

import (
	"fmt"
	"github.com/MTthoas/API-Github-GO/functions"
	"github.com/joho/godotenv"
	"log"
	
	
)

func main() {
	fmt.Println("Hello, World!")

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
	
}
