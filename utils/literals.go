package utils

import (
	"unicode"
)

func RunesIsDigit(runes []rune) bool {
	if runes == nil {
		return false
	}
	for _, char := range runes {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}
