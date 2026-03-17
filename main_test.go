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

// These tests validate the HTTP server contract described in prd.md.
// Keep each case small and focused (Red → Green → Clean workflow).

// TestHomePage checks the GET / handler renders the required HTML form fields.
func TestHomePage(t *testing.T) {
	// Arrange: build a plain GET / request.
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	// Act: call handler directly (no network server needed).
	homeHandler(rec, req)

	// Assert: request succeeded and page contains form + required fields.
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

// TestAsciiArtPostValid verifies valid text and banner return HTTP 200 and expected ASCII output.
func TestAsciiArtPostValid(t *testing.T) {
	// Arrange: valid text and explicit banner.
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

	// Build expected ASCII output using the same converter pipeline.
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

// TestAsciiArtPostEmptyTextReturnsBadRequest ensures empty submissions are rejected (400).
func TestAsciiArtPostEmptyTextReturnsBadRequest(t *testing.T) {
	// Arrange: submit empty text; this is invalid for conversion.
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

// TestAsciiArtPostInvalidBannerReturnsNotFound ensures missing banner files return 404.
func TestAsciiArtPostInvalidBannerReturnsNotFound(t *testing.T) {
	// Arrange: unknown banner value should trigger missing-banner path.
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

// TestAsciiArtPostInvalidCharacterReturnsBadRequest validates non-printable ASCII characters are rejected.
func TestAsciiArtPostInvalidCharacterReturnsBadRequest(t *testing.T) {
	// Arrange: include tab, which is outside supported ASCII range.
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

// TestAsciiArtPostMissingTemplateReturnsNotFound confirms missing template produces 404 response.
func TestAsciiArtPostMissingTemplateReturnsNotFound(t *testing.T) {
	// Arrange: temporarily move template file away to simulate deployment misconfiguration.
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

// TestAsciiArtPostMissingBannerFileReturnsNotFound confirms missing banner file returns 404.
func TestAsciiArtPostMissingBannerFileReturnsNotFound(t *testing.T) {
	// Arrange: temporarily hide one banner file and post matching banner name.
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
