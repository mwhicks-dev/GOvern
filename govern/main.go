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
		text = Clean(text)
		if LowerCase(text) == LowerCase("ENTER") {
			// TODO: Use enter workflow
		} else if LowerCase(text) == LowerCase("CLOSE") {
			break
		} else if LowerCase(text) == LowerCase("SIGNUP") {
			// TODO: Use signup workflow
		} else if text == "" {
			// Do nothing
		} else {
			fmt.Printf("Did not recognize command %v.\n", text)
		}
	}

}
