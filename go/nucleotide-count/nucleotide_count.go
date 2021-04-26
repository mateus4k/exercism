package dna

import (
	"fmt"
	"strings"
)

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, nucleotide := range strings.ToUpper(string(d)) {
		_, finded := h[nucleotide]

		if !finded {
			return nil, fmt.Errorf("invalid nucleotide %v", nucleotide)
		}

		h[nucleotide]++
	}

	return h, nil
}
