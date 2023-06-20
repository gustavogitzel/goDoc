package utils

import (
	"strings"
	"unicode"
)

// CleanNonDigits removes non-digit characters from the provided string pointer doc
func CleanNonDigits(doc *string) {
	var result strings.Builder
	for _, character := range *doc {
		if unicode.IsDigit(character) {
			result.WriteRune(character)
		}
	}

	*doc = result.String()
}

// IsAllDigits checks whether all characters in the provided string pointer doc are digits.
func IsAllDigits(doc *string) bool {
	for _, char := range *doc {
		if !unicode.IsDigit(char) {
			return false
		}
	}

	return true
}

// ToInt converts a Unicode character representing a digit to its corresponding integer value.
func ToInt(character rune) int {
	return int(character - '0')
}
