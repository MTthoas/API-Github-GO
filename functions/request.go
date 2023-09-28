package functions

import (
	"io/ioutil"
	"net/http"
	"github.com/MTthoas/API-Github-GO/model"
	"github.com/gofiber/fiber/v2/log"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)


func GetMethod(url, token string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("User-Agent", "github-repo-list")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

func FetchRepos(user string, repos []model.Repository) {

	var wg sync.WaitGroup

	for _, repo := range repos {
		wg.Add(1)

		go func(r model.Repository) {
			defer wg.Done()

			repoURL := fmt.Sprintf("https://github.com/%s/%s.git", user, r.Name)
			if err := CloneRepository(repoURL, "./repos/"+r.Name); err != nil {
				log.Error("Failed to clone %s: %s", r.Name, err)
				return
			}

			if err := GitPullAndFetch("./repos/" + r.Name); err != nil {
				log.Error("Failed to update %s: %s", r.Name, err)
				return
			}
		}(repo)
	}

	wg.Wait()      // Attendre que toutes les goroutines soient terminées
}

func CloneRepository(repoURL, targetDir string) error {

	// Vérifier si le dossier cible existe déjà
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		// Le dossier n'existe pas, alors nous allons le créer
		if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
			return fmt.Errorf("erreur lors de la création du dossier cible: %v", err)
		}
	}

	cmd := exec.Command("git", "clone", repoURL, targetDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func GitPullAndFetch(repoDir string) error {

	/* Déplacement, attribution à Currentbranch, pull & fetch heads*/

	// Get the current branch
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = repoDir
	currentBranch, err := cmd.Output()
	if err != nil {
		return err
	}
	branchName := strings.TrimSpace(string(currentBranch))

	// Switch to the current branch
	cmd = exec.Command("git", "checkout", branchName)
	cmd.Dir = repoDir
	if err := cmd.Run(); err != nil {
		return err
	}

	// Pull the latest changes
	cmd = exec.Command("git", "pull")
	cmd.Dir = repoDir
	if err := cmd.Run(); err != nil {
		return err
	}

	// Fetch all branches
	cmd = exec.Command("git", "fetch", "--all")
	cmd.Dir = repoDir
	return cmd.Run()
}

