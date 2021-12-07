package github

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Repo struct {
	Name        string `json:"name"`
	Project     string `json:"project"`
	Description string `json:"description"`
}

func GetRepos() ([]Repo, error) {

	resp, err := http.Get("http://localhost:9017/github/projects")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var repos []Repo
	err = json.Unmarshal(body, &repos)
	if err != nil {
		return nil, err
	}

	return repos, nil
}
















type Client struct {}

func (s *Client) GetRepos() ([]Repo, error) {
	return GetRepos()
}

