package main

import "fmt"

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

func main() {
	x := []string{"a", "b", "c"}
	y := []string{"x", "y", "z"}
	z := []string{"a", "b", "c"}

	fmt.Println(equal(x, y), equal(x, z))
}