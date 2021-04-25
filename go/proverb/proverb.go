package proverb

import "fmt"

const (
	stanza = "For want of a %s the %s was lost."
	last   = "And all for the want of a %s."
)

func Proverb(rhyme []string) []string {
	result := make([]string, len(rhyme))

	for i := range rhyme {
		if i < len(rhyme)-1 {
			result[i] = fmt.Sprintf(stanza, rhyme[i], rhyme[i+1])
		} else {
			result[i] = fmt.Sprintf(last, rhyme[0])
		}
	}

	return result
}
