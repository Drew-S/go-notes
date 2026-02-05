package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// TODO: Add data storage location option [optional]

type Shortcut struct {
	Command  string `yaml:"command"`  // command to use when `notes new {command}`
	Template string `yaml:"template"` // ./template.md -> {config.Template}/template.md
	Location string `yaml:"location"` // ./inbox ->
}

type Config struct {
	Root      string     `yaml:"root"`
	Templates string     `yaml:"templates"` // ./templates/ -> {config.Root}/templates/
	Home      string     `yaml:"home"`
	Shortcuts []Shortcut `yamls:"shortcuts"`
}

func pathTransform(input, dot_slash, home string) string {
	if strings.HasPrefix(input, "$HOME") || strings.HasPrefix(input, "~") {
		input = filepath.Join(home, strings.TrimPrefix(strings.TrimPrefix(input, "$HOME"), "~"))
	}

	if strings.HasPrefix(input, "./") {
		input = filepath.Join(dot_slash, strings.TrimPrefix(input, "./"))
	}

	return input
}

func LoadConfig() (Config, error) {
	HOME := os.Getenv("HOME")

	X := os.Getenv("XDG_CONFIG_HOME")

	locations := []string{
		fmt.Sprintf("%s/go-notes/go-notes.yaml", X),
		fmt.Sprintf("%s/go-notes/config.yaml", X),
		fmt.Sprintf("%s/.config/go-notes/go-notes.yaml", HOME),
		fmt.Sprintf("%s/.config/go-notes/config.yaml", HOME),
		fmt.Sprintf("%s/.go-notes/go-notes.yaml", HOME),
		fmt.Sprintf("%s/.go-notes/config.yaml", HOME),
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

	if cfg.Root == "" {
		log.Fatal("missing fields: ensure that root is set")
		return Config{}, err
	}

	// Expand config to be absolute path
	cfg.Root = pathTransform(cfg.Root, "", HOME)

	cfg.Templates = pathTransform(cfg.Templates, cfg.Root, HOME)

	for i := range cfg.Shortcuts {
		cfg.Shortcuts[i].Template = pathTransform(cfg.Shortcuts[i].Template, cfg.Templates, HOME)
		cfg.Shortcuts[i].Location = pathTransform(cfg.Shortcuts[i].Location, cfg.Root, HOME)
	}

	return cfg, nil
}
