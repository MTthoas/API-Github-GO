package functions

import (
	"os"
	"path/filepath"
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