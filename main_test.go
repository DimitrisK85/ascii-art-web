package main

import (
	"ascii-art-web/internal/banner"
	"ascii-art-web/internal/converter"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHomePage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	homeHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	body := rec.Body.String()
	if !strings.Contains(body, `<form action="/ascii-art" method="POST">`) {
		t.Fatalf("expected HTML form to post to /ascii-art, got body: %q", body)
	}
	if !strings.Contains(body, `name="text"`) {
		t.Fatalf("expected text field in form, got body: %q", body)
	}
	if !strings.Contains(body, `name="banner"`) {
		t.Fatalf("expected banner selector in form, got body: %q", body)
	}
}

func TestAsciiArtPostValid(t *testing.T) {
	form := url.Values{}
	form.Add("text", "Hi")
	form.Add("banner", "standard")

	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	asciiArtHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	charMap, err := banner.LoadBannerFile(filepath.Join("banners", "standard.txt"))
	if err != nil {
		t.Fatalf("load banner failed: %v", err)
	}
	expected := strings.Join(converter.ConvertText(charMap, "Hi"), "\n")
	body := rec.Body.String()
	if !strings.Contains(body, expected) {
		t.Fatalf("expected converted output %q in response body, got %q", expected, body)
	}
}

func TestAsciiArtPostEmptyTextReturnsBadRequest(t *testing.T) {
	form := url.Values{}
	form.Add("text", "")
	form.Add("banner", "standard")

	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	asciiArtHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestAsciiArtPostInvalidBannerReturnsNotFound(t *testing.T) {
	form := url.Values{}
	form.Add("text", "Hello")
	form.Add("banner", "invalid-banner")

	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	asciiArtHandler(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}

func TestAsciiArtPostInvalidCharacterReturnsBadRequest(t *testing.T) {
	form := url.Values{}
	form.Add("text", "Hello\tWorld")
	form.Add("banner", "standard")

	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	asciiArtHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
	body := rec.Body.String()
	if !strings.Contains(body, "400 Bad Request: unsupported characters") {
		t.Fatalf("expected detailed bad request reason, got %q", body)
	}
}

func TestAsciiArtPostMissingTemplateReturnsNotFound(t *testing.T) {
	templatePath := filepath.Join("templates", "index.html")
	backupPath := templatePath + ".bak"
	if err := os.Rename(templatePath, backupPath); err != nil {
		t.Fatalf("failed to move template for test: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Rename(backupPath, templatePath); err != nil {
			t.Fatalf("failed to restore template: %v", err)
		}
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	homeHandler(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
	body := rec.Body.String()
	if !strings.Contains(body, "404 Not Found") {
		t.Fatalf("expected not found message in response body, got %q", body)
	}
}

func TestAsciiArtPostMissingBannerFileReturnsNotFound(t *testing.T) {
	bannerPath := filepath.Join("banners", "thinkertoy.txt")
	backupPath := bannerPath + ".bak"
	if err := os.Rename(bannerPath, backupPath); err != nil {
		t.Fatalf("failed to move banner for test: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Rename(backupPath, bannerPath); err != nil {
			t.Fatalf("failed to restore banner: %v", err)
		}
	})

	form := url.Values{}
	form.Add("text", "Hello")
	form.Add("banner", "thinkertoy")

	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	asciiArtHandler(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
	body := rec.Body.String()
	if !strings.Contains(body, "404 Not Found") {
		t.Fatalf("expected not found message in response body, got %q", body)
	}
}
