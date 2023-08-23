package capitalize_string
import "strings"

func CapitalizeString(string_to_capitalize string) string {
	return strings.ToUpper(string_to_capitalize[0:1]) + string_to_capitalize[1:len(string_to_capitalize)]
}
