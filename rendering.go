package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func renderTemplate(w http.ResponseWriter, name string, data pageData) error {
	tmplPath := filepath.Join(templateDir, name)
	if _, err := os.Stat(tmplPath); err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	return tmpl.Execute(w, data)
}

func renderNotFound(w http.ResponseWriter, data pageData, message string) {
	if data.Banner == "" {
		data.Banner = "standard"
	}
	data.Error = "404 Not Found: " + message
	w.WriteHeader(http.StatusNotFound)
	if err := renderTemplate(w, "index.html", data); err != nil {
		handleTemplateError(w, err)
	}
}

func renderBadRequest(w http.ResponseWriter, data pageData, message string) {
	if data.Banner == "" {
		data.Banner = "standard"
	}
	data.Error = "400 Bad Request: " + message
	w.WriteHeader(http.StatusBadRequest)
	if err := renderTemplate(w, "index.html", data); err != nil {
		handleTemplateError(w, err)
	}
}

func handleTemplateError(w http.ResponseWriter, err error) {
	if os.IsNotExist(err) {
		http.Error(w, "404 Not Found: template missing", http.StatusNotFound)
		return
	}

	http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
}
