package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		text = Clean(text)
		fmt.Println("<" + text + ">")
	}

}

func Clean(s string) string {

	// strip whitespace
	s = strings.TrimSpace(s)

	// remove duplicate whitespace
	for {
		s0 := s
		s = strings.Replace(s, "  ", " ", -1)
		s = strings.Replace(s, "\t", " ", -1)
		s = strings.Replace(s, "\n", " ", -1)
		s = strings.Replace(s, "\v", " ", -1)
		s = strings.Replace(s, "\f", " ", -1)
		s = strings.Replace(s, "\r", " ", -1)

		if strings.Compare(s, s0) == 0 {
			break
		}
	}

	return s

}
