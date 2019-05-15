package main

import "fmt"
import "bufio"
import "os"
import "log"

func shell() string {
	var first string
	fmt.Scanln(&first)
	return first
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		//	fmt.Println("data is being piped to stdin")
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	} else {
		for {
			fmt.Println(shell())
		}
	}

}