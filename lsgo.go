package main

import (
	"io/ioutil"
	"log"

	"github.com/fatih/color"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for index, file := range files {
		blue := color.New(color.FgHiBlue)
		yellow := color.New(color.FgYellow)
		if index == len(files)-1 {
			if file.IsDir() {
				yellow.Print(" D: " + file.Name())
			} else {
				blue.Print(" F: " + file.Name())
			}
		} else {
			if file.IsDir() {
				yellow.Println(" D: " + file.Name())
			} else {
				blue.Println(" F: " + file.Name())
			}
		}
	}
}
