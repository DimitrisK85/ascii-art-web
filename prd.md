# Product Requirements Document (PRD)
## ASCII-Art Web

### Project Overview
A Go web application that converts user text into ASCII art using predefined banner templates.

### Objectives
- Convert input text into ASCII art using a selected banner.
- Support banners: `standard`, `shadow`, `thinkertoy`.
- Render the result in the browser through HTML templates.
- Keep behavior simple and aligned with the existing conversion logic.

### Technical Requirements

#### 1. HTTP Endpoints
- `GET /`
  - Returns the main HTML page.
  - Contains a text input and a banner selector.
  - Contains a submit control that sends a `POST` to `/ascii-art`.
- `POST /ascii-art`
  - Receives user text and selected banner from form data.
  - Converts the text to ASCII art.
  - Returns an HTML response that displays the result.
  - The result may be shown on the same page or a dedicated `/ascii-art` page.

#### 2. Banner Support
- Allowed banners: `standard`, `shadow`, `thinkertoy`.
- Banner files are loaded from:
  - `banners/standard.txt`
  - `banners/shadow.txt`
  - `banners/thinkertoy.txt`
- Unknown banners are not selectable in the normal UI flow; if the selected banner file is missing, return `404 Not Found`.

#### 3. Conversion Behavior
- The conversion core is the existing `internal/banner` loader and `internal/converter` logic.
- Input text must be ASCII 32..126 and newline (`\n`) for multiline handling, if already supported by converter behavior.
- Rendered output is returned as monospaced preformatted text (for example using `<pre>`).
- No additional transformation rules beyond current conversion behavior.

#### 4. HTML Templates
- Templates must be located in a root-level `templates/` directory.
- At minimum:
  - `templates/index.html`
  - optionally `templates/ascii-art.html` if result is rendered on a separate page.

#### 5. Status Codes
- `200 OK` for successful requests.
- `400 Bad Request` for invalid form data (missing text, unsupported characters, malformed submissions).
- `404 Not Found` for missing required resources (templates or banner files).
- `500 Internal Server Error` for unexpected runtime failures.

#### 6. Implementation Constraints
- Server must be written in Go.
- Use only Go standard library packages.
- Keep code readable, explicit, and minimal.
- Apply test-first workflow for new behavior.

### Development Workflow
- Add/adjust tests before implementation (TDD).
- Keep changes focused and aligned with this PRD.
- Do not introduce features outside this PRD unless explicitly requested.

### Success Criteria
- `GET /` returns the form UI.
- Submitting valid data to `POST /ascii-art` returns ASCII-art output.
- Submitting invalid request data (empty text, unsupported characters, malformed submissions) returns clear `400`.
- Missing template/banner conditions are handled with `404`.
- Unsupported runtime failures return `500`.
