package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bsdpunk/SimpleGoShell/shell/bget"
)

func Shell() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		words := strings.Fields(text)
		fmt.Print(words, len(words))
		switch {
		case words[0] == "quit":
			os.Exit(3)
		case words[0] == "bget":
			fmt.Println(bget.Web(words[1]))
		default:
			fmt.Println(words)
		}
	}

	return "exit"
}
