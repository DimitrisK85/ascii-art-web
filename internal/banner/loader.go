package banner

import (
	"bufio"
	"fmt"
	"os"
)

// LoadBannerFile reads a banner file and returns a map of runes to their ASCII art representation.
// Each character is represented by 8 lines of ASCII art.
// Characters are separated by empty lines in the banner file.
func LoadBannerFile(input string) (output map[rune][]string, err error) {
	// Open the banner file
	file, err := os.Open(input)
	if err != nil {
		return output, fmt.Errorf("failed to open banner file: %w", err)
	}
	defer file.Close()

	asciiMap := make(map[rune][]string)
	scanner := bufio.NewScanner(file)

	// Start from ASCII 32 (space character)
	currentChar := rune(32)
	var block []string

	for scanner.Scan() {
		line := scanner.Text()

		// Empty line separates characters (each character has 8 lines)
		if line == "" && len(block) == 8 {
			asciiMap[currentChar] = block
			block = []string{}
			currentChar++
		} else if line != "" || len(block) > 0 {
			// Add line to current character block
			block = append(block, line)
		}
	}

	// Handle last character if file doesn't end with separator
	if len(block) == 8 {
		asciiMap[currentChar] = block
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return asciiMap, nil
}
