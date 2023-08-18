package main
import "fmt"

func main(){
    string_to_count := "I really think I am getting the hang of Go now."

    fmt.Printf("The String \"%s\" is %d characters long\n", string_to_count, len(string_to_count))
}
