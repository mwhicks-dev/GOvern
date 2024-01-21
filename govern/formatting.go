package main

import "strings"

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

func LowerCase(s string) string {

	return strings.ToLower(s)

}
