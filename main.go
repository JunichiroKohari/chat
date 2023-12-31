package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type template1Handler struct {
	once      sync.Once
	filename  string
	template1 *template.Template
}

func (t *template1Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.template1 = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.template1.Execute(w, nil)
}

func main() {
	http.Handle("/", &template1Handler{filename: "chat.html"})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
