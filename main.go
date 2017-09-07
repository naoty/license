package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

// Version is the version of this application. This value is set by Makefile.
var Version = "0.0.0"

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v", "--version":
			fmt.Println(Version)
			os.Exit(0)
		}
	}

	path := filepath.Join("templates", "mit.txt")
	template, err := template.ParseFiles(path)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	year := time.Now().Year()
	fullname, err := getFullnameFromGit()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	data := map[string]interface{}{
		"year":     year,
		"fullname": fullname,
	}

	template.Execute(os.Stdout, data)
}

func getFullnameFromGit() (string, error) {
	cmd := exec.Command("git", "config", "user.name")
	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%s", out)
	result = strings.TrimRight(result, "\n")

	return result, nil
}
