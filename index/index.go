package index

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

type Post struct {
	ID   int
	Name string
	Body string
}

type TemplateData struct {
	Posts []Post
}

func fetchAllPosts(db *sql.DB) ([]Post, error) {
	posts := make([]Post, 0)
	rows, err := db.Query("select * from posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.ID, &p.Name, &p.Body)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	// XXX is this necessary?
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

type IndexHandler struct {
	Tmpl *template.Template
	DB   *sql.DB
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	posts, err := fetchAllPosts(h.DB)
	if err != nil {
		log.Fatal(err)
	}
	data := TemplateData{Posts: posts}
	if err := h.Tmpl.ExecuteTemplate(w, "template.txt", data); err != nil {
		log.Fatal(err)
	}
}

func New(db *sql.DB) (*IndexHandler, error) {
	tmpl, err := template.ParseFiles("template.txt")
	if err != nil {
		return nil, err
	}
	return &IndexHandler{Tmpl: tmpl, DB: db}, nil
}
