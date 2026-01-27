package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
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

	switch command {
	case "diary":
	case "journal":
		_, err := CreateDirIfNotExist(config.Journal)
		if err != nil {
			panic(err)
		}
		OpenNvim(config.Root, filepath.Join(config.Journal, fmt.Sprintf("%s.md", time.Now().Format("2006-01-02"))))
		break

	case "inbox":
	case "new":
		_, err := CreateDirIfNotExist(config.Inbox)
		if err != nil {
			panic(err)
		}
		OpenNvim(config.Root, filepath.Join(config.Inbox, fmt.Sprintf("%s.md", time.Now().Format("2006-01-02"))))
		break

	default:
	case "default":
		_, err := CreateDirIfNotExist(config.Home)
		if err != nil {
			panic(err)
		}
		OpenNvim(config.Root, config.Home)
	}
}
