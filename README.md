# ASCII-Art-Web

## Description
ASCII-Art-Web is a simple Go HTTP server that converts text to ASCII art using predefined banners.

It provides a browser form at `/` where users enter text and choose a banner, then submits to `/ascii-art`.

## Authors
- Project contributors: Dimitris Kolias

## Usage: how to run
- Start the server:
  - `go run .`
- Open `http://localhost:8080` in a browser.
- Enter text, select one of:
  - `standard`
  - `shadow`
  - `thinkertoy`
- Submit the form to see ASCII output on the same page.

## Implementation details: algorithm
- Route handlers:
  - `GET /`: renders `templates/index.html`.
  - `POST /ascii-art`: reads form fields `text` and `banner`.
- Validation:
  - request method is checked per route.
  - input must be non-empty.
  - input characters are limited to ASCII printable range (`32..126`) and newline (`\n`).
- Conversion:
  - newline normalization handles `\r\n` and `\r` so browser text is consistent.
  - load the selected banner file from `banners/<banner>.txt`.
  - use `internal/banner` to parse banner data.
  - use `internal/converter` to build ASCII lines.
  - join generated lines with `\n` and render inside `<pre>`.
- Error rendering:
  - `400 Bad Request`: invalid form/missing text/unsupported input.
  - `404 Not Found`: missing `templates/index.html` or selected banner file.
  - `500 Internal Server Error`: unexpected render failures.
- Project uses only Go standard library packages.

## Notes
- Banner is selected from a fixed UI dropdown, so unsupported banner values are not expected in normal browser usage.
- `normalizeInput` removes Windows/newline artifacts (`\r\n`, `\r`) so form submissions behave consistently across OSs.

## Instructions
- Endpoints implemented:
  - `GET /`
  - `POST /ascii-art`
- Templates are located in `templates/`.
- Banner files are in `banners/`.
- Required Go modules are standard library only.
