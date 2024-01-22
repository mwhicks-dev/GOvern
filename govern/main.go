package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Please enter ENTER, SIGNUP, or CLOSE (case-insensitive): ")
		text, _ := reader.ReadString('\n')

		Handle(reader, Clean(text))
		fmt.Println()
	}

}

func Handle(reader *bufio.Reader, cmd string) {

	if LowerCase(cmd) == LowerCase("ENTER") {
		EnterLoop(reader)
	} else if LowerCase(cmd) == LowerCase("CLOSE") {
		os.Exit(0)
	} else if LowerCase(cmd) == LowerCase("SIGNUP") {
		SignupLoop(reader)
	} else if cmd == "" {
		// Do nothing
	} else {
		fmt.Printf("Did not recognize command %v.\n", cmd)
	}

}
