package main
import "fmt"

func main(){
	string_to_reverse := "I like these exercises."
	reversed_string := ""

	for current_index := len(string_to_reverse) - 1; current_index >= 0; current_index-- {
		reversed_string += string(string_to_reverse[current_index])
	}
	fmt.Print(reversed_string)
}
