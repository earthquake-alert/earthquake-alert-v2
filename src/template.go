package src

import (
	"html/template"
	"path/filepath"
	"strings"
)

var funcmap = template.FuncMap{
	"convert": Convert,
}

const FILE_PATH = "../templates"

func Template(fileName string, obj any) (string, error) {
	path := filepath.Join(FILE_PATH, fileName)

	templ, err := template.New(fileName).Funcs(funcmap).ParseFiles(path)
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
