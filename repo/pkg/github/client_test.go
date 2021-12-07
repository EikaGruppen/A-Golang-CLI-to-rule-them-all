package github_test

import (
	"fmt"
	"repo/pkg/github"
	"testing"
)

func TestGetRepos(t *testing.T) {
	repos, err := github.GetRepos()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", repos)
	// TODO asserts
}
