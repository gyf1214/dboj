package view

import (
	"html/template"
	"io"
	"path/filepath"
)

const (
	root   = "layout"
	layout = "layout.html"
	dir    = "view"
)

func parseTemplate(name string) func(io.Writer, interface{}) error {
	lp := filepath.Join(dir, layout)
	fp := filepath.Join(dir, name)
	tplt := template.Must(template.ParseFiles(lp, fp))
	return func(w io.Writer, d interface{}) error {
		return tplt.ExecuteTemplate(w, root, d)
	}
}

// templates
var (
	Index = parseTemplate("index.html")
	Login = parseTemplate("login.html")
)
