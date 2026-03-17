package converter

import (
	"ascii-art-web/internal/banner"
	"reflect"
	"strings"
	"testing"
)

func TestConvertLine(t *testing.T) {
	charMap, err := banner.LoadBannerFile("../../banners/standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}

	// Test 1: "Hello"
	t.Run("Hello", func(t *testing.T) {
		input := "Hello"
		expected := []string{
			" _    _          _   _          ",
			"| |  | |        | | | |         ",
			"| |__| |   ___  | | | |   ___   ",
			"|  __  |  / _ \\ | | | |  / _ \\  ",
			"| |  | | |  __/ | | | | | (_) | ",
			"|_|  |_|  \\___| |_| |_|  \\___/  ",
			"                                ",
			"                                ",
		}

		result := ConvertLine(charMap, input)

		if len(result) != 8 {
			t.Errorf("Expected 8 lines, got %d", len(result))
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected:\n%v\n\nGot:\n%v", expected, result)
		}
	})

	// Test 2: "123"
	t.Run("123", func(t *testing.T) {
		input := "123"
		result := ConvertLine(charMap, input)

		if len(result) != 8 {
			t.Errorf("Expected 8 lines, got %d", len(result))
		}
	})

	// Test 3: "A B" (with space)
	t.Run("A B", func(t *testing.T) {
		input := "A B"
		result := ConvertLine(charMap, input)

		if len(result) != 8 {
			t.Errorf("Expected 8 lines, got %d", len(result))
		}

		// Verify space is handled (should have gap between A and B)
		if len(result[0]) == 0 {
			t.Error("Result should not be empty")
		}
	})
}

func TestConvertText(t *testing.T) {
	charMap, err := banner.LoadBannerFile("../../banners/standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}

	t.Run("Single line", func(t *testing.T) {
		input := "Hi"
		result := ConvertText(charMap, input)
		if len(result) != 8 {
			t.Errorf("Expected 8 lines, got %d", len(result))
		}
	})

	t.Run("Multiple lines", func(t *testing.T) {
		input := "Hello\nWorld"
		result := ConvertText(charMap, input)
		if len(result) != 16 {
			t.Errorf("Expected 16 lines (8 per line), got %d", len(result))
		}
	})

	t.Run("Empty line", func(t *testing.T) {
		input := "Hello\n\nWorld"
		result := ConvertText(charMap, input)
		if len(result) != 17 {
			t.Errorf("Expected 17 lines, got %d", len(result))
		}
	})
}

func TestConvertTextWithColor_SubstringMatch(t *testing.T) {
	charMap, err := banner.LoadBannerFile("../../banners/standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}

	result := ConvertTextWithColor(charMap, "kitten kit", "kit", "\033[31m")

	if len(result) != 8 {
		t.Fatalf("Expected 8 lines, got %d", len(result))
	}

	if result[0] == "" {
		t.Fatal("Expected non-empty first line")
	}

	if !strings.Contains(result[0], "\033[31m") || !strings.Contains(result[0], "\033[0m") {
		t.Fatalf("Expected ANSI color markers in output line, got: %q", result[0])
	}
}

func TestConvertTextWithColor_WholeStringWhenSubstringEmpty(t *testing.T) {
	charMap, err := banner.LoadBannerFile("../../banners/standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}

	plain := ConvertText(charMap, "A")
	colored := ConvertTextWithColor(charMap, "A", "", "\033[31m")

	if len(colored) != 8 {
		t.Fatalf("Expected 8 lines, got %d", len(colored))
	}

	for i := 0; i < 8; i++ {
		expected := "\033[31m" + plain[i] + "\033[0m"
		if colored[i] != expected {
			t.Fatalf("Expected %q, got %q", expected, colored[i])
		}
	}
}
