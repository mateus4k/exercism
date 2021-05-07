package reverse

func Reverse(input string) string {
	output := make([]rune, 0, len(input))
	runes := []rune(input)

	for i := len(runes) - 1; i >= 0; i-- {
		output = append(output, runes[i])
	}

	return string(output)
}
