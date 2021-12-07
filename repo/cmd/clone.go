package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)

func CloneRepo(rs repoStorage, stdin io.ReadCloser, stdout io.WriteCloser) {
	repos, err := rs.GetRepos()
	if err != nil {
		fmt.Println(err)
		return
	}

	choosen := prompt(repos, stdin, stdout)
	clone(choosen)
}

func clone(repo Repo) {
	fmt.Printf("Cloning: %s...", repo.Name)
	time.Sleep(2 * time.Second) // TODO actually do the clone command
	fmt.Println("Done!")
}

func prompt(repos []Repo, stdin io.ReadCloser, stdout io.WriteCloser) (choosen Repo) {

	templates := &promptui.SelectTemplates{
		Label:    `Choose repo`,
		Active:   "{{ .Name | cyan }} ({{ .Project }})",
		Inactive: "{{ .Name | cyan | faint }} ({{ .Project }})",
		Selected: " ",
		Details: `
--------- Description ----------
{{ .Description }}`,
	}

	prompt := promptui.Select{
		Items:     repos,
		Templates: templates,
		Searcher: func(input string, index int) bool {
			return strings.Contains(repos[index].Name, input)
		},
		StartInSearchMode: true,
		HideSelected:      true,
		Stdin: stdin,
		Stdout: stdout,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return repos[i]
}
