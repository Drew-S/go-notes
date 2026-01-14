package main

import (
	"os"
	"os/exec"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	command := "default"

	if len(os.Args) >= 2 {
		command = os.Args[2]
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if err = os.Chdir(config.Root); err != nil {
		panic(err)
	}

	switch command {
	case "journal":
		break
	case "new":
		break
	default:
	case "default":
		nvim := exec.Command("nvim", config.Home)
		nvim.Stdin = os.Stdin
		nvim.Stdout = os.Stdout
		nvim.Stderr = os.Stderr
		if err = nvim.Run(); err != nil {
			panic(err)
		}
	}

	if err := os.Chdir(cwd); err != nil {
		panic(err)
	}
}
