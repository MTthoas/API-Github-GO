package functions

import (
	"os"
	"path/filepath"
	"sort"
	"github.com/MTthoas/API-Github-GO/model"
	"strconv"
	"encoding/csv"
	"github.com/gofiber/fiber/v2/log"
)


func RemoveContents(dir string) error {

	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		entryPath := filepath.Join(dir, entry.Name())

		// Généralement on manipule des dossiers, mais on peut aussi avoir des fichiers ( au cas où )

		if entry.IsDir() {
			// Si c'est un répertoire, récursivement supprimer son contenu
			if err := os.RemoveAll(entryPath); err != nil {
				return err
			}
		} else {
			// Si c'est un fichier, le supprimer
			if err := os.Remove(entryPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func SortByDate(repos []model.Repository) []model.Repository {
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].UpdatedAt > repos[j].UpdatedAt
	})
	return repos
}

func ExcelWrite(repos []model.Repository , writer *csv.Writer) {

	// Specification lié à la structure de mes Headers

	headers := []string{
		"Name", "Full Name", "HTML URL", "Description", "Language",
		"Created At", "Updated At", "Stargazers Count", "Forks Count",
		"Watchers Count", "Open Issues Count", "Default Branch",
		"Owner Login", "Owner Avatar URL", "Owner HTML URL",
	}
	
	if err := writer.Write(headers); err != nil {
		log.Fatalf("Failed to write headers to CSV: %s", err)
	}
	
	// Write the repositories' data to the CSV file
	for _, repo := range repos {
		row := []string{
			repo.Name,
			repo.FullName,
			repo.HTMLURL,
			repo.Description,
			repo.Language,
			repo.CreatedAt,
			repo.UpdatedAt,
			strconv.Itoa(repo.StargazersCount),
			strconv.Itoa(repo.ForksCount),
			strconv.Itoa(repo.WatchersCount),
			strconv.Itoa(repo.OpenIssuesCount),
			repo.DefaultBranch,
			repo.Owner.Login,
			repo.Owner.AvatarURL,
			repo.Owner.HTMLURL,
		}
		if err := writer.Write(row); err != nil {
			log.Fatalf("Failed to write repository data to CSV: %s", err)
		}else{
			log.Info("Ecriture des données en cours...", repo.Name,)  
		}
	}

	log.Info("Ecriture des données terminée")
}
