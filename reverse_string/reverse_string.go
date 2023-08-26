package reverse_string

func ReverseString(string_to_reverse string) string {
	runes_to_reverse := []rune(string_to_reverse)
	for index := 0; index < len(runes_to_reverse) / 2; index++ {
		runes_to_reverse[index], runes_to_reverse[len(runes_to_reverse) - 1 - index] = runes_to_reverse[len(runes_to_reverse) - 1 - index], runes_to_reverse[index]
	}
	return string(runes_to_reverse)
}
