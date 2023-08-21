package letter_in_string

import "strings"

func LetterInString(string_to_search_in string, substring_to_search string) bool {
	return strings.Contains(string_to_search_in, substring_to_search)
}