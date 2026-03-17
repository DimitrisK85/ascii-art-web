package main

import (
	"log"
	"net/http"
)

const (
	serverAddress = ":8080"
	templateDir   = "templates"
)

type pageData struct {
	Input  string
	Banner string
	Result string
	Error  string
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	log.Printf("Starting server at http://localhost%s", serverAddress)
	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		log.Fatal(err)
	}
}
