package count_string
import "fmt"

func CountString(string_to_count string) int {
	fmt.Printf("The String \"%s\" is %d characters long\n", string_to_count, len(string_to_count))
	return len(string_to_count)
}