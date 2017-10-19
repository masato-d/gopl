package main

import "fmt"

func rotate(s []int, i int) []int {
	rtn := s[i:]
	rtn = append(rtn, s[:i]...)
	return rtn
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	r := rotate(s, 2)
	fmt.Println(r)
}