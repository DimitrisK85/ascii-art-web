package main

import (
	"ascii-art-web/internal/banner"
	"ascii-art-web/internal/converter"
	"net/http"
	"path/filepath"
	"strings"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderBadRequest(w, pageData{Banner: "standard"}, "wrong method, expected GET")
		return
	}

	data := pageData{Banner: "standard"}
	if err := renderTemplate(w, "index.html", data); err != nil {
		handleTemplateError(w, err)
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		renderBadRequest(w, pageData{Banner: "standard"}, "wrong method, expected POST")
		return
	}

	if err := r.ParseForm(); err != nil {
		renderBadRequest(w, pageData{Banner: "standard"}, "invalid form submission")
		return
	}

	input := r.FormValue("text")
	input = normalizeInput(input)
	bannerName := r.FormValue("banner")
	if bannerName == "" {
		bannerName = "standard"
	}

	if input == "" {
		renderBadRequest(w, pageData{Input: input, Banner: bannerName}, "empty text")
		return
	}

	if !isValidAsciiInput(input) {
		renderBadRequest(w, pageData{Input: input, Banner: bannerName}, "unsupported characters")
		return
	}

	bannerPath := filepath.Join("banners", bannerName+".txt")
	charMap, err := banner.LoadBannerFile(bannerPath)
	if err != nil {
		renderNotFound(w, pageData{Input: input, Banner: bannerName}, "Banner not found")
		return
	}

	art := converter.ConvertText(charMap, input)
	result := strings.Join(art, "\n")

	data := pageData{Input: input, Banner: bannerName, Result: result}
	if err := renderTemplate(w, "index.html", data); err != nil {
		handleTemplateError(w, err)
	}
}
