package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Function to create parent directory if it does not currently
// exist, otherwise does nothing
// Returns true if created and false if not
func CreateDirIfNotExist(file string) (bool, error) {
	dir := filepath.Dir(file)
	dir_created := false

	err := os.MkdirAll(dir, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatalln(err)
		return false, err
	}

	if os.IsExist(err) {
		dir_created = true
	}

	return dir_created, nil
}

// file is absolute path
// root is notes directory, also absolute
func OpenNvim(root, file string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = os.Chdir(root)
	if err != nil {
		return err
	}

	nvim := exec.Command("nvim", file)
	nvim.Stdin = os.Stdin
	nvim.Stdout = os.Stdout
	nvim.Stderr = os.Stderr

	err = nvim.Run()
	if err != nil {
		return err
	}

	err = os.Chdir(cwd)
	if err != nil {
		return err
	}

	return nil
}
