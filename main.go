package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	path := filepath.Join("templates", "mit.txt")
	template, err := template.ParseFiles(path)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	data := map[string]string{
		"year":     "2017",
		"fullname": "Naoto Kaneko",
	}

	template.Execute(os.Stdout, data)
}
