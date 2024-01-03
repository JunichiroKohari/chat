package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/JunichiroKohari/go_trace"
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
	t.template1.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "app address")
	flag.Parse()
	r := newRoom()
	r.tracer = go_trace.New(os.Stdout)
	http.Handle("/", &template1Handler{filename: "chat.html"})
	http.Handle("/room", r)
	go r.run()
	log.Println("Start Web Server. Port: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
