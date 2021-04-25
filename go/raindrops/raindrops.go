package raindrops

import "strconv"

func Convert(amount int) string {
	var remainder string

	var translations = []struct {
		Number int
		Word   string
	}{
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"},
	}

	for _, translation := range translations {
		if amount%translation.Number == 0 {
			remainder += translation.Word
		}
	}

	if remainder == "" {
		remainder = strconv.Itoa(amount)
	}

	return remainder
}
