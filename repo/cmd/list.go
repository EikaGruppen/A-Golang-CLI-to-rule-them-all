package cmd

import (
	"fmt"
	"repo/pkg/github"
)

func ListRepos() {
	repos, err := github.GetRepos()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, repo := range repos {
		fmt.Printf("%s - %s\n\t%s\n\n", green(repo.Project), green(repo.Name), repo.Description)
	}

}

func green(text string) string {
	return fmt.Sprintf("\033[1;32m%s\033[0m", text)
}

