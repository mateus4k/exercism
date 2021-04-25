package hamming

import "errors"

func Distance(a string, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("length is not equal")
	}

	var count int

	for index, item := range []byte(a) {
		if item != b[index] {
			count++
		}
	}

	return count, nil
}
