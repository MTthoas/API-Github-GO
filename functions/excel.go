package functions

import (
	"os"
	"path/filepath"
	"sort"
	"github.com/MTthoas/API-Github-GO/model"
	"strconv"
	"encoding/csv"
	"github.com/gofiber/fiber/v2/log"
	"archive/zip"
	"io"
	"strings"
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
		} 
	}

	return nil
}

func ManageCSV(csvFileName string, repos []model.Repository) {
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

    ExcelWrite(repos, writer)
    log.Info("Fichier CSV créé")
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
	
	// Ecriture dans le fichier CSV des données de chaque repository, lié à la structure de mes Headers et de mes données ( model/repository.go )

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

func ArchiveRepositories(sourceDir string, destinationZip string) error {
	zipFile, err := os.Create(destinationZip)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()


	// Walk est une fonction récursive qui parcourt le dossier sourceDir et ses sous-dossiers

	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create header for the file or directory
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = filepath.Join(filepath.Base(sourceDir), strings.TrimPrefix(path, sourceDir))

		if info.IsDir() {
			_, err = zipWriter.CreateHeader(header)

			if err != nil {
				return err
			}

		} else {

			// L'utilisation de CreateHeader permet de créer un fichier avec les bons droits
			// ( par exemple, les fichiers exécutables seront créés avec les droits d'exécution )

			writer, err := zipWriter.CreateHeader(header)
			if err != nil {
				return err
			}
			fileToArchive, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fileToArchive.Close()

			_, err = io.Copy(writer, fileToArchive)
			if err != nil {
				return err
			}
		}
		return nil
	})
}



