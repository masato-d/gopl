package main

import "fmt"

func filter(strings []string) []string {
	n := len(strings)
	i := 1
	for i < n {
		if strings[i-1] == strings[i] {
			copy(strings[i-1:], strings[i:])
			n--
		} else {
			i++
		}
	}
	return strings[:n]
}

func main() {
	s := []string{"a", "b", "b", "c", "a", "c", "d", "d", "d", "e"}
	fmt.Println(filter(s))
}