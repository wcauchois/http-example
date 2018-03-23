package index

import (
	"html/template"
	"log"
	"net/http"
)

type TemplateData struct {
	Name string
}

type IndexHandler struct {
	Tmpl *template.Template
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{Name: "bill"}
	if err := h.Tmpl.ExecuteTemplate(w, "template.txt", data); err != nil {
		log.Fatal(err)
	}
}

func New() (*IndexHandler, error) {
	tmpl, err := template.ParseFiles("template.txt")
	if err != nil {
		return nil, err
	}
	return &IndexHandler{tmpl}, nil
}
