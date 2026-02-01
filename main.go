package main

import (
	"flag"
	"fmt"
)

/*
TODO:
	- data storage
		- sqlite indexes for tags, links, etc.
    - tag system
		- update tag index from files, frontmatter and inline (async)
		- search for tags
		- associate via tags
	- link system
		- update forward/backward links from files (async)
		- generate link graph
			- page data for template use
	- bookmark system
		- set, add, remove bookmarks
		- quick open bookmark
			- parameterize quick open? (current date journal, newest note,
			  oldest note, first/last via tag, etc.)
    - richer template system
        - more variables to work with
		- input variables (so when called, it asks for user input)
			- user label, type, and where it goes
			- whether this is the filename expansion template
			- or the template itself
				- example {{input:int:age}} {{input:str:source}} {{input:sel[book,video,site]:url}}
					- {{input:str:source}} {{input::source}} --> {{.Input.Source}} (filename expansion)
					- then all {{.Input.Source}} gets the template data gets this
				- how to reuse (for example source is used twice, we don't
				  want the user to input twice)
	- search system using fzf (filename, infile, in frontmatter, in content,
      via tag, via bookmark, etc.)
*/

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
				executeTemplate(short.Template, short.Location)
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
