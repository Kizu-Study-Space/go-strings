package capitalize_string
import "strings"

func CapitalizeString(string_to_capitalize string) string {
	capitalized_string := strings.ToUpper(string(string_to_capitalize[0]))
	capitalized_string += strings.TrimPrefix(string_to_capitalize, string(string_to_capitalize[0]))
	return capitalized_string
}
