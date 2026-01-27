package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"time"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		panic(err)
	}

	flag.Parse()

	command := flag.Arg(0)

	switch command {
	case "new":
		shortcut := flag.Arg(1)
		if shortcut == "" {
			fmt.Println("No command given")
			return
		}
		for _, short := range config.Shortcuts {
			if short.Command == shortcut {
				_, err := CreateDirIfNotExist(short.Location)
				if err != nil {
					panic(err)
				}
				OpenNvim(config.Root, filepath.Join(
					short.Location,
					fmt.Sprintf("%s.md", time.Now().Format("2006-01-02"))))
				return
			}
		}

	default:
		_, err := CreateDirIfNotExist(config.Home)
		if err != nil {
			panic(err)
		}
		OpenNvim(config.Root, config.Home)
	}
}
