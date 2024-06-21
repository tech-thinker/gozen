package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
)


func GenerateCode(tplFile string, data interface{}) (string, error) {
	tpl, err := template.ParseFiles(tplFile)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func WriteFile(path, data string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Fprintf(file, `%s`, data)
	return nil
}

func CreateDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
