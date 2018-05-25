package view

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
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

func handleStatic(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if p := recover(); p != nil {
			http.Error(w, fmt.Sprint(p), 400)
		}
	}()
	lp, err := filepath.Rel("/", r.URL.Path)
	if err != nil {
		panic(err)
	}
	fp := filepath.Join(dir, layout)
	tplt, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}
	err = tplt.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		panic(err)
	}
}

func ServeStatic() {
	http.HandleFunc("/view/", handleStatic)
}

// templates
var (
	Index = parseTemplate("index.html")
	Login = parseTemplate("login.html")
)
