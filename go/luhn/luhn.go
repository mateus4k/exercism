package luhn

import (
	"strings"
	"unicode"
)

func Valid(s string) bool {
	var sum int
	var tag bool = false

	s = strings.ReplaceAll(s, " ", "")

	if len(s) <= 1 {
		return false
	}

	for i := len(s) - 1; i >= 0; i-- {
		r := rune(s[i])

		if !unicode.IsDigit(r) {
			return false
		}

		n := int(r - '0')

		if tag {
			n *= 2
			if n > 9 {
				sum -= 9
			}
		}

		sum += n
		tag = !tag
	}

	return sum%10 == 0
}
