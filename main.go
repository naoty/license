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

func main() {
	path := filepath.Join("templates", "mit.txt")
	template, err := template.ParseFiles(path)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
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
