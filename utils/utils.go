package utils

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"os"
)


func GenerateCode(tplFS embed.FS, tplFile string, data interface{}) (string, error) {
	tpl, err := template.ParseFS(tplFS, tplFile)
	if err != nil {
		return "{}", err
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return "{}", err
	}

	return buf.String(), nil
}

func WriteFile(path, data string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
    _, err = fmt.Fprintf(file, `%s`, data)
    if err != nil {
        return err
    }
	return nil
}

func CreateDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
