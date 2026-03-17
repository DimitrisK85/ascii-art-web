package main

import (
	"ascii-art-web/internal/banner"
	"ascii-art-web/internal/converter"
	"net/http"
	"path/filepath"
	"strings"
)

// homeHandler serves the main page using GET only.
// It renders the form and keeps the default banner selected.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// Reject unsupported methods with 400 and an explicit error message.
		renderBadRequest(w, pageData{Banner: "standard"}, "wrong method, expected GET")
		return
	}

	// Default view is rendered with standard banner selected.
	data := pageData{Banner: "standard"}
	if err := renderTemplate(w, "index.html", data); err != nil {
		handleTemplateError(w, err)
	}
}

// asciiArtHandler receives user text and banner selection, validates input,
// converts text to ASCII art, and renders the same page with the result.
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// This route only supports POST submissions.
		renderBadRequest(w, pageData{Banner: "standard"}, "wrong method, expected POST")
		return
	}

	if err := r.ParseForm(); err != nil {
		// Form parsing failed (malformed request body).
		renderBadRequest(w, pageData{Banner: "standard"}, "invalid form submission")
		return
	}

	// Read submitted data.
	input := r.FormValue("text")
	input = normalizeInput(input)
	bannerName := r.FormValue("banner")

	// Default banner matches the previous CLI behavior and UI default.
	if bannerName == "" {
		bannerName = "standard"
	}

	if input == "" {
		// Empty payload is invalid input for conversion.
		renderBadRequest(w, pageData{Input: input, Banner: bannerName}, "empty text")
		return
	}

	// Reject characters outside printable ASCII and newline.
	if !isValidAsciiInput(input) {
		renderBadRequest(w, pageData{Input: input, Banner: bannerName}, "unsupported characters")
		return
	}

	// Build path to the chosen banner and load its character map.
	bannerPath := filepath.Join("banners", bannerName+".txt")
	charMap, err := banner.LoadBannerFile(bannerPath)
	if err != nil {
		// Unknown/missing banner file is treated as 404.
		renderNotFound(w, pageData{Input: input, Banner: bannerName}, "Banner not found")
		return
	}

	// Convert whole text into ASCII art lines and join by newline for rendering.
	art := converter.ConvertText(charMap, input)
	result := strings.Join(art, "\n")

	data := pageData{Input: input, Banner: bannerName, Result: result}
	if err := renderTemplate(w, "index.html", data); err != nil {
		handleTemplateError(w, err)
	}
}
