package main

import (
	"io/ioutil"
	"log"

	"github.com/fatih/color"
)

func FormatMessage(s string, id bool) string {
	var outputMessage string
	if id {
		outputMessage = " D: "
	} else {
		outputMessage = " F: "
	}
	return outputMessage + s
}

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
				yellow.Print(FormatMessage(file.Name(), file.IsDir()))
			} else {
				blue.Print(FormatMessage(file.Name(), file.IsDir()))
			}
		} else {
			if file.IsDir() {
				yellow.Println(FormatMessage(file.Name(), file.IsDir()))
			} else {
				blue.Println(FormatMessage(file.Name(), file.IsDir()))
			}
		}
	}
}
