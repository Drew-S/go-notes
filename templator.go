package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

type Data struct {
	Time time.Time
}

func NewData() Data {
	return Data{
		Time: time.Now(),
	}
}

func executeTemplate(template_location, output_location string) error {

	// If the file exists, do not run the execution of the template
	_, err := os.Stat(output_location)
	if err != nil && !os.IsNotExist(err) {
		log.Fatalln(err)
		return err
	}
	if !os.IsNotExist(err) {
		return nil
	}

	// Execute template
	content, err := os.ReadFile(template_location)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	data := NewData()

	t := template.Must(template.New("file").Parse(string(content)))

	file, err := os.Create(output_location)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	defer file.Close()

	t.Execute(file, data)
	return nil
}
