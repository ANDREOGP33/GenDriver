package view

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func ParseFS(fs fs.FS, pattern ...string) Template {
	t, err := template.ParseFS(fs, pattern...)
	if err != nil {
		log.Fatal(err.Error())
	}
	return Template{
		htmlTpl: t,
	}
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Fatal(err.Error())
	}
}
