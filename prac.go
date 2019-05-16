package main

import "fmt"
import "strings"

func main() {
	someString := "one two three four "
	
	words := strings.Fields(someString)
	
	fmt.Println(words, len(words))
	fmt.Println(words[1])
}
