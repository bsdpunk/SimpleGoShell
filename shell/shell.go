package shell

import "fmt"
import "os"
import "strings"

func Shell() string {
	var command string
	fmt.Print("$ ")
	fmt.Scanln(&command)
	words := strings.Fields(command)
	switch {
	case words[0] == "quit":
		os.Exit(3)
	default:

		return command
	}
	return command
}
