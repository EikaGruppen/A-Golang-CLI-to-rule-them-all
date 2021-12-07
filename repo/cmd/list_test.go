package cmd_test

import (
	"bytes"
	"repo/cmd"
	"strings"
	"testing"
)

var alfa = cmd.Repo{
	Name:        "alfa",
	Project:     "ap",
	Description: "ad",
}

var beta = cmd.Repo{
	Name:        "beta",
	Project:     "bp",
	Description: "bd",
}

type reposMock struct{}

func (m *reposMock) GetRepos() ([]cmd.Repo, error) {
	return []cmd.Repo{alfa, beta}, nil
}

func TestListRepos(t *testing.T) {
	mock := &reposMock{}

	var b bytes.Buffer
	capturer := promptCapture{&b}

	cmd.ListRepos2(mock, &capturer)

	output := capturer.String()
	if !strings.Contains(output, alfa.Name) {
		t.Fatal("Missing 'alfa'!")
	}
}

type promptCapture struct {
	*bytes.Buffer
}

func (m *promptCapture) Close() error {
	return nil
}
