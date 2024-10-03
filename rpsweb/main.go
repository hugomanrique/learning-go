package main

import (
	"fmt"
	"net/http"
	"rps/handlers"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	port := "8080"
	fmt.Println("Serving on port", port)
	http.ListenAndServe(":"+port, router)
}
