package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename,p.Body,0644)
}

func load(title string) (*Page,error){
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title:title,Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hola me encantan los %s","Hugo")
}


func main(){
	// p1 := &Page{Title: "Hugo", Body: []byte("Hugo is a static site generator written in Go.")}
	// p1.save()
	
	// p2, _ := load("Hugo")
	// fmt.Println(string(p2.Body))

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil) )
}