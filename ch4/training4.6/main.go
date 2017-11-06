package main

import (
	"fmt"
	"unicode"
)

func compress_space(str []byte) []byte {
	i := 0
	for _, s := range str {
		if unicode.IsSpace(rune(s)) && unicode.IsSpace(rune(str[i-1])) {
			continue
		}
		str[i] = s
		i++
	}
	return str[:i]
}

func main() {
	str := []byte("abc d e  fg h    i jk l   mn")
	fmt.Println(string(compress_space(str)))
}