package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func main() {
	fmt.Println("hello world")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tm, err := template.Must(template.ParseFiles(filepath.Join(cwd, "../pages/index.html")))
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
