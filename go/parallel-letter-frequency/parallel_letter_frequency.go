package letter

// FreqMap records the frequency of each rune in a gven text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency uses concurrency to calculate the frequency of letters
func ConcurrentFrequency(words []string) FreqMap {
	ch := make(chan FreqMap, len(words))
	fm := FreqMap{}

	for _, s := range words {
		go func(s string) {
			ch <- Frequency(s)
		}(s)
	}

	for range words {
		for k, v := range <-ch {
			fm[k] += v
		}
	}

	return fm
}
