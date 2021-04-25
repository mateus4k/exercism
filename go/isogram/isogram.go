// Package isogram detects isograms
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns a boolean indicating whether the word is an isogram or not
func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for _, r := range word {
		if !unicode.IsLetter(r) {
			continue
		}

		if strings.Count(word, string(r)) > 1 {
			return false
		}
	}

	return true
}
