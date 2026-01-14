package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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

	// Load config
	// TODO: Add in multiple format support (check multiple locations)
	//       - ~/.config/go-notes/config.yaml
	//       - ~/.config/go-notes/go-notes.yaml
	//       - ~/.go-notes.yaml
	config_dir := os.Getenv("XDG_CONFIG_HOME") + "/go-notes/"
	if config_dir == "/go-notes/" {
		config_dir = "/.config/go-notes/"
	}

	config_dir = HOME + "/" + config_dir

	data, err := os.ReadFile(fmt.Sprintf("%sconfig.yaml", config_dir))
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

	// Expand config to be absolute path
	cfg.Root = strings.Replace(cfg.Root, "$HOME", HOME, 1)
	cfg.Root = strings.Replace(cfg.Root, "~", HOME, 1)

	cfg.Templates = strings.Replace(cfg.Templates, "./", cfg.Root+"/", 1)
	cfg.Journal = strings.Replace(cfg.Journal, "./", cfg.Root+"/", 1)
	cfg.Inbox = strings.Replace(cfg.Inbox, "./", cfg.Root+"/", 1)

	return cfg, nil
}

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
