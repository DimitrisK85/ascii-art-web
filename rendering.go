package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// renderTemplate loads and executes an HTML template with the provided page data.
// It returns parsing/execution errors for caller-side handling.
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

// renderNotFound renders the home page with a 404 error context.
// Used when a required resource (e.g., template/banner file) is missing.
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

// renderBadRequest renders the home page with a 400 error context.
// Used for invalid methods, input, or malformed form submissions.
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

// handleTemplateError maps template/load failures to the required HTTP status:
// missing files are reported as 404, all other failures as 500.
func handleTemplateError(w http.ResponseWriter, err error) {
	if os.IsNotExist(err) {
		http.Error(w, "404 Not Found: template missing", http.StatusNotFound)
		return
	}

	http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
}
