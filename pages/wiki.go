package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0644)
}

func load(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := load(title)
	t, _ := template.ParseFiles("view.html")
	t.Execute(w, p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	fmt.Printf(title)
	p, _ := load(title)

	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

func main() {
	// p1 := &Page{Title: "Hugo", Body: []byte("Hugo is a static site generator written in Go.")}
	// p1.save()

	// p2, _ := load("Hugo")
	// fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", handler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
