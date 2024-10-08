package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"rps/rps"
	"strconv"
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
	restartValue()
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
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)
	out, error := json.MarshalIndent(result, "", "  ")
	if error != nil {
		log.Println(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", nil)
}

func restartValue() {
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0
}
