package utils

import (
	"bytes"
	"embed"
	"html/template"
	"strings"
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

func ApplyEscapeChar(data string) string {
	data = strings.ReplaceAll(data, "&lt;", "<")
	data = strings.ReplaceAll(data, "&gt;", ">")
	data = strings.ReplaceAll(data, "&amp;", "&")
	return data
}
