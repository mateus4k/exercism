// Package isogram detects isograms
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns a boolean indicating whether the word is an isogram or not
func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for index, letter := range word {
		if unicode.IsLetter(letter) && strings.ContainsRune(word[index+1:], letter) {
			return false
		}
	}

	return true
}
