package main

import (
	"fmt"
	"strings"
)

func main() {
	const success = "the two strings is an anagram"
	const fail = "thw two strings is NOT an anagram"
	s1 := "stop"
	s2 := "pots"

	if len(s1) != len(s2) {
		fmt.Println(fail)
		return
	}

	checked := s2
	for i := 0; i < len(s1) ; i++ {
		if !strings.Contains(checked, string(s1[i])) {
			fmt.Println(fail)
			return
		}
		checked = strings.Replace(checked, string(s1[i]), "", 1)
	}

	if len(checked) > 0 {
		fmt.Println(fail)
		return
	}

	fmt.Println(success)
}