package main
import "fmt"

func main(){
	string_to_reverse := "I like these exercises."
	reversed_string := ""
	counter := len(string_to_reverse)

	for 0 < counter {
		counter--
		reversed_string += string(string_to_reverse[counter])
	}
	fmt.Print(reversed_string)
}