package main

import "strings"

// normalizeInput standardizes line endings for all platforms:
// Windows CRLF -> LF, then removes remaining CR characters.
func normalizeInput(input string) string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.ReplaceAll(input, "\r", "")
	return input
}

// isValidAsciiInput verifies each input rune is printable ASCII (32-126)
// or newline. Any other rune is rejected.
func isValidAsciiInput(input string) bool {
	for _, char := range input {
		if char == '\n' {
			continue
		}
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}
