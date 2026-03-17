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

// ConvertText converts text with real newline characters to ASCII art.
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

// ConvertTextWithColor converts text to ASCII art and colors matched characters.
// If substring is empty, all input characters are colored.
func ConvertTextWithColor(charMap map[rune][]string, input, substring, colorCode string) []string {
	lines := strings.Split(input, "\n")
	var result []string

	for _, line := range lines {
		converted := convertLineWithColor(charMap, line, substring, colorCode)
		result = append(result, converted...)
	}

	return result
}

// convertLineWithColor appends each glyph row and wraps matched characters with ANSI color/reset.
func convertLineWithColor(charMap map[rune][]string, text, substring, colorCode string) []string {
	output := make([]string, 8)
	if text == "" {
		return []string{""}
	}

	colorMask := buildColorMask(text, substring)
	resetCode := "\033[0m"

	for i, char := range text {
		artLines := charMap[char]
		shouldColor := i < len(colorMask) && colorMask[i]
		for row := 0; row < 8; row++ {
			if shouldColor {
				output[row] += colorCode + artLines[row] + resetCode
			} else {
				output[row] += artLines[row]
			}
		}
	}

	return output
}

// buildColorMask marks character indices that belong to substring matches in text.
func buildColorMask(text, substring string) []bool {
	mask := make([]bool, len(text))
	if substring == "" {
		for i := range mask {
			mask[i] = true
		}
		return mask
	}

	offset := 0
	for {
		idx := strings.Index(text[offset:], substring)
		if idx == -1 {
			break
		}
		start := offset + idx
		end := start + len(substring)
		for i := start; i < end && i < len(mask); i++ {
			mask[i] = true
		}
		offset = end
	}

	return mask
}
