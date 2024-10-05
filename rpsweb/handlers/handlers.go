package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	templateDir  = "templates/"
	templateBase = templateDir + "base.html"
)

func renderTemplate(w http.ResponseWriter, page string, data any) {
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))
	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func Index(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "index.html", nil)
	// tpl, err := template.ParseFiles("templates/base.html", "templates/index.html")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// data := struct {
	// 	Title   string
	// 	Message string
	// }{
	// 	Title:   "Pagina de inicio",
	// 	Message: "Bienvenido a nuestra pagina de inicio",
	// }

	// err = tpl.Execute(w, data)

}

type Player struct {
	Name string
}

var player Player

func NewGame(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "new-game.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		player.Name = r.Form.Get("name")
	}

	if player.Name == "" {
		renderTemplate(w, "new-game.html", nil)
		return
	}
	renderTemplate(w, "game.html", player)
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Play")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", nil)
}
