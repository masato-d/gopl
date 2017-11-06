package main

import "fmt"

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	str := []byte("reverse")
	fmt.Println(str)
	reverse(str)
	fmt.Println(str)
}