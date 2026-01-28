package main

import (
	"flag"
	"fmt"
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
				short.Location = expandFileName(short.Location)
				_, err := CreateDirIfNotExist(short.Location)
				if err != nil {
					panic(err)
				}
				OpenNvim(config.Root, short.Location)
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
