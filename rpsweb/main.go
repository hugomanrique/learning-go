package main

import (
	"fmt"
	"net/http"
	"rps/handlers"
)

func main() {

	router := http.NewServeMux()

	// Manage static files
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)
	port := "8080"
	fmt.Println("Serving on port", port)
	http.ListenAndServe(":"+port, router)
}
