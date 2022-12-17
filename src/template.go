package src

import (
	"html/template"
	"path/filepath"
	"strings"
)

const FILE_PATH = "../templates"

func Template(fileName string, obj any) (string, error) {
	path := filepath.Join(FILE_PATH, fileName)

	templ, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}

	writer := new(strings.Builder)

	err = templ.Execute(writer, obj)
	if err != nil {
		return "", err
	}

	return writer.String(), nil
}
