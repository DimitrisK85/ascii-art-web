package converter

import "strings"

// ConvertLine converts a single line of text to ASCII art.
// Returns 8 lines of ASCII art (or empty string "" for empty input).
func ConvertLine(charMap map[rune][]string, text string) []string {
	// Create 8 empty lines for ASCII art output
	output := make([]string, 8)
	// Handle empty string case
	if text == "" {
		emptyOutput := []string{""}
		return emptyOutput
	}
	// Process each character in the input text
	for _, char := range text {
		// Get the 8-line ASCII art representation for this character
		artLines := charMap[char]

		// Append each line of the character's art to the corresponding output line
		for i := 0; i < 8; i++ {
			output[i] += artLines[i]
		}
	}

	return output
}

// ConvertText converts multi-line input into a single ASCII-art slice.
// Splits input by \n and converts each line separately.
func ConvertText(char map[rune][]string, input string) []string {
	// Split input by real newlines only.
	lines := strings.Split(input, "\n")
	var result []string

	// Convert each line and append to result
	for _, line := range lines {
		convertedLine := ConvertLine(char, line)
		result = append(result, convertedLine...)
	}

	return result
}
