package reverse_string

func ReverseString(string_to_reverse string) string {
	reversed_string := ""

	for current_index := len(string_to_reverse) - 1; current_index >= 0; current_index-- {
		reversed_string += string(string_to_reverse[current_index])
	}

	return reversed_string
}
