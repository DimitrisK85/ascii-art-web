package main

import (
	"log"
	"net/http"
)

const (
	// serverAddress is the local address where the HTTP server listens.
	serverAddress = ":8080"
	// templateDir stores the directory that contains all HTML templates.
	templateDir   = "templates"
)

// pageData is the payload passed to the HTML template.
// It carries user input, selected banner, generated ASCII result, and an error message.
type pageData struct {
	Input  string
	Banner string
	Result string
	Error  string
}

func main() {
	// Register route handlers for the required endpoints.
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	// Start the web server.
	log.Printf("Starting server at http://localhost%s", serverAddress)
	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		log.Fatal(err)
	}
}
