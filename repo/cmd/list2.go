package cmd

import (
	"fmt"
	"io"
)

func ListRepos2(rs repoStorage, w io.WriteCloser) {
	repos, err := rs.GetRepos()
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	for _, repo := range repos {
		fmt.Fprintf(w, "%s - %s\n\t%s\n\n", green(repo.Project), green(repo.Name), repo.Description)
	}
}
