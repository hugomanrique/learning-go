package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "New Game")
}

func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Game")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Play")
}
