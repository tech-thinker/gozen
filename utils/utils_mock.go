package utils

import "embed"

//go:embed *.tmpl
var tmplFS embed.FS

func GetMockTmpl() embed.FS {
	return tmplFS
}
