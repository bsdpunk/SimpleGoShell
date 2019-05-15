package shell

import "fmt"

func Shell() string {
	var first string
	fmt.Print("$ ")
	fmt.Scanln(&first)
	return first
}
