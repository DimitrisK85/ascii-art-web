package main

import "strings"

func normalizeInput(input string) string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.ReplaceAll(input, "\r", "")
	return input
}

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
