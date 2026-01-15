package main

import (
	"log"
	"net/http"
	"viewer/internal/handlers"
)

func main() {
	// Serve Static Files (CSS, JS)
	// This tells GO: if a URL starts with /static/, look inside the web/static folder"
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Register Routes. When someone visits "/", run the HomeHandler function
	http.HandleFunc("/", handlers.HomeHandler)

	// Starting the server
	log.Println("Starting server on http://localhost:8080")
	// This starts the server and blocks forever. If it crashes it returns an error.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

