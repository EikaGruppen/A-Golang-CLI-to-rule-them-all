package cmd

import "repo/pkg/github"

type Repo = github.Repo

type repoStorage interface {
	GetRepos() ([]Repo, error)
}
