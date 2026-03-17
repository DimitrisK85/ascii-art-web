package banner

import "testing"

func TestLoadBannerFile(t *testing.T) {
	input := "../../banners/standard.txt"

	result, err := LoadBannerFile(input)
	if err != nil {
		t.Errorf("LoadBannerFile(%s) error = %v", input, err)
	}

	if result == nil {
		t.Error("Expected character map, but got nil")
	}

	if _, exists := result[' ']; !exists {
		t.Error("Expected space character in map")
	}

	if lines, exists := result[' ']; exists {
		if len(lines) != 8 {
			t.Errorf("Expected 8 lines for space character, got %d", len(lines))
		}
	}
}

func TestLoadBannerFile_MissingFile(t *testing.T) {
	// Test 2: Error handling for missing file
	input := "nonexistent.txt"

	_, err := LoadBannerFile(input)

	// Check that error is returned
	if err == nil {
		t.Error("Expected error for missing file, but got nil")
	}
}
