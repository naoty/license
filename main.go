package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

func main() {
	path := filepath.Join("templates", "mit.txt")
	template, err := template.ParseFiles(path)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	year := time.Now().Year()

	data := map[string]interface{}{
		"year":     year,
		"fullname": "Naoto Kaneko",
	}

	template.Execute(os.Stdout, data)
}
