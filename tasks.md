# ascii-art-web Task Cards

## TASK-01 PRD Alignment (Web Scope)
- Finalize requirements only for web flow: home page + ASCII conversion by banner.
- Keep only shared behavior already implemented in `internal/banner` and `internal/converter`.

## TASK-02 HTTP Server Foundation
- Implement HTTP server in `main.go`.
- Register `GET /` and `POST /ascii-art` routes.
- Return plain text/HTML error messages with correct status codes.

## TASK-03 Templates and Routing UI
- Create `templates/index.html` with:
  - text input
  - banner selector (`standard`, `shadow`, `thinkertoy`)
  - submit action posting to `/ascii-art`
- Add result view either in `/ascii-art` response page or same route.
- Keep all templates in root-level `templates/` folder.

## TASK-04 Conversion Integration
- Use existing banner loader and converter for processing.
- Validate banner input and request payload.
- Convert text and render output in `<pre>`.

## TASK-05 Error Handling and HTTP Status
- `200 OK` on successful page renders and conversion.
- `400 Bad Request` for invalid form data (missing text, unsupported banner).
- `404 Not Found` if required template or banner file is unavailable.
- `500 Internal Server Error` for unhandled processing/render failures.

## TASK-06 Handler Tests (TDD)
- Add `main_test.go`-style handler tests with `net/http/httptest`.
- Cover:
  - `GET /` success with form fields.
  - `POST /ascii-art` success.
  - invalid banner and invalid request data.
  - missing template path behavior.

## TASK-07 Documentation
- Create/update `README.MD` with required sections:
  - Description
  - Authors
  - Usage: how to run
  - Implementation details: algorithm
  - Instructions

## TASK-08 Clean-up and Consistency
- Remove unused CLI-only code only if clearly outside the new project scope.
- Keep conversion core files unless they conflict with runtime goals.
- Ensure only standard library packages are used.
