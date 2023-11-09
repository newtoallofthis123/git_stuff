package main

import (
	"fmt"
	"os"
	"os/exec"
)

const VERSION string = "0.0.1"

var allowed map[string]bool = map[string]bool{
	"help":    true,
	"version": true,
	"i":       true,
	"o":       true,
	"c":       true,
}

func main() {
	fmt.Println("Qit")

	if !isGitRepo() {
		fmt.Println("Not a git repository")
		return
	}

	// parse command line arguments
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No arguments given")
		return
	}

	switch args[0] {
	case "version":
		version()
		return
	case "help":
		help()
		return
	case "o":
		gitUrlOpen()
		return
	default:
		help()
	}
}

func version() {
	fmt.Println("Version " + VERSION)
}

func isGitRepo() bool {
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		return false
	}
	return true
}

func help() {
	fmt.Println("Help")
	fmt.Println("Here are the allowed ones: ")
	for arg := range allowed {
		fmt.Print(arg + " ")
	}
	fmt.Println()
}

func gitUrlOpen() {
	//get the url from the git config
	output, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
	if err != nil {
		fmt.Println("Does this repo have a remote origin?")
		return
	}
	url := string(output)

	err = exec.Command("open", url).Run()
	if err != nil {
		fmt.Println("Error opening url: " + url)
		return
	}
}
