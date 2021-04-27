package accumulate

func Accumulate(given []string, converter func(string) string) []string {
	res := make([]string, len(given))

	for i, v := range given {
		res[i] = converter(v)
	}

	return res
}
