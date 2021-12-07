package main

import (
	"repo/cmd"
	"repo/pkg/github"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		println("Please specify subcommand")
		os.Exit(1)
	}

	subCmd := args[0]

	switch subCmd {
	case "list":
		cmd.ListRepos()
	case "clone":
		cmd.CloneRepo(&github.Client{}, os.Stdin, os.Stdout)
	default:
		println("Unknown subcommand")
		os.Exit(1)
	}

}
