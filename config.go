package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Root      string `yaml:"root"`
	Templates string `yaml:"templates"`
	Journal   string `yaml:"journal"`
	Inbox     string `yaml:"inbox"`
	Home      string `yaml:"home"`
}

func LoadConfig() (Config, error) {
	HOME := os.Getenv("HOME")

	X := os.Getenv("XDG_CONFIG_HOME")

	locations := []string{
		fmt.Sprintf("%s/go-notes/go-notes.yaml", X),
		fmt.Sprintf("%s/go-notes/config.yaml", X),
		fmt.Sprintf("%s/.config/go-notes/go-notes.yaml", HOME),
		fmt.Sprintf("%s/.config/go-notes/config.yaml", HOME),
		"./.go-notes.yaml",
	}

	config_file := ""

	for _, loc := range locations {
		_, err := os.Stat(loc)
		if err == nil {
			config_file = loc
			break
		}
	}

	data, err := os.ReadFile(config_file)
	if err != nil {
		log.Fatal("Unable to read config file, did you create it?")
		return Config{}, err
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal("Unable to parse config file, ensure its formatted correct")
		return Config{}, err
	}

	if cfg.Root == "" || cfg.Inbox == "" || cfg.Journal == "" {
		log.Fatal("missing fields: ensure that root, inbox, and journal are set")
		return Config{}, err
	}

	// Expand config to be absolute path
	cfg.Root = strings.Replace(cfg.Root, "$HOME", HOME, 1)
	cfg.Root = strings.Replace(cfg.Root, "~", HOME, 1)

	cfg.Templates = strings.Replace(cfg.Templates, "./", cfg.Root+"/", 1)
	cfg.Templates = strings.Replace(cfg.Templates, "$HOME", HOME, 1)
	cfg.Templates = strings.Replace(cfg.Templates, "~", HOME, 1)

	cfg.Journal = strings.Replace(cfg.Journal, "./", cfg.Root+"/", 1)
	cfg.Journal = strings.Replace(cfg.Journal, "$HOME", HOME, 1)
	cfg.Journal = strings.Replace(cfg.Journal, "~", HOME, 1)

	cfg.Inbox = strings.Replace(cfg.Inbox, "./", cfg.Root+"/", 1)
	cfg.Inbox = strings.Replace(cfg.Inbox, "$HOME", HOME, 1)
	cfg.Inbox = strings.Replace(cfg.Inbox, "~", HOME, 1)

	return cfg, nil
}
