package gopdf

import (
	"fmt"
)

// Helper functions
func SplitIntoLines(text string) []string {
	var lines []string
	currentLine := ""

	for _, char := range text {
		if char == '\n' || char == '\r' {
			if len(currentLine) > 0 {
				lines = append(lines, currentLine)
				currentLine = ""
			}
		} else {
			currentLine += string(char)
		}
	}

	if len(currentLine) > 0 {
		lines = append(lines, currentLine)
	}

	return lines
}

func FindVietnameseLines(lines []string) map[int]string {
	vietnameseChars := "àáạảãâầấậẩẫăằắặẳẵèéẹẻẽêềếệểễìíịỉĩòóọỏõôồốộổỗơờớợởỡùúụủũưừứựửữỳýỵỷỹđĐ"
	result := make(map[int]string)

	for i, line := range lines {
		for _, char := range vietnameseChars {
			if containsRune(line, char) {
				result[i+1] = line
				break
			}
		}
	}

	return result
}

func containsRune(text string, r rune) bool {
	for _, char := range text {
		if char == r {
			return true
		}
	}
	return false
}

func AnalyzeCharacters(text string) {
	fmt.Printf("📊 Character Analysis:\n")

	charCount := make(map[rune]int)
	totalChars := 0

	for _, char := range text {
		charCount[char]++
		totalChars++
	}

	fmt.Printf("  Total characters: %d\n", totalChars)
	fmt.Printf("  Unique characters: %d\n", len(charCount))

	// Show first 20 unique characters
	fmt.Printf("  Sample characters: ")
	count := 0
	for char := range charCount {
		if count >= 20 {
			break
		}
		if char >= 32 && char <= 126 { // Printable ASCII
			fmt.Printf("'%c' ", char)
		} else {
			fmt.Printf("U+%04X ", char)
		}
		count++
	}
	fmt.Println()

	// Check for common Vietnamese characters
	vietnameseChars := []rune{'à', 'á', 'ạ', 'ả', 'ã', 'â', 'ầ', 'ấ', 'ậ', 'ẩ', 'ẫ', 'ă', 'ằ', 'ắ', 'ặ', 'ẳ', 'ẵ', 'đ', 'Đ'}
	vietnameseFound := false
	for _, vnChar := range vietnameseChars {
		if charCount[vnChar] > 0 {
			fmt.Printf("  Vietnamese char '%c': %d times\n", vnChar, charCount[vnChar])
			vietnameseFound = true
		}
	}

	if !vietnameseFound {
		fmt.Printf("  ❌ No Vietnamese characters detected\n")
	}
}

func FindTextInContent(content, searchTerm string) bool {
	// Case insensitive search
	contentLower := ""
	searchLower := ""

	for _, char := range content {
		if char >= 'A' && char <= 'Z' {
			contentLower += string(char + 32) // Convert to lowercase
		} else {
			contentLower += string(char)
		}
	}

	for _, char := range searchTerm {
		if char >= 'A' && char <= 'Z' {
			searchLower += string(char + 32)
		} else {
			searchLower += string(char)
		}
	}

	// Simple contains check
	return containsSubstring(contentLower, searchLower)
}

func containsSubstring(text, substring string) bool {
	if len(substring) == 0 {
		return true
	}
	if len(text) < len(substring) {
		return false
	}

	for i := 0; i <= len(text)-len(substring); i++ {
		match := true
		for j := 0; j < len(substring); j++ {
			if text[i+j] != substring[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}

	return false
}
