package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	categories := make(map[string]int)
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
		case unicode.IsLetter(r):
			categories["L"]++
		case unicode.IsMark(r):
			categories["M"]++
		case unicode.IsNumber(r):
			categories["N"]++
		case unicode.IsPunct(r):
			categories["P"]++
		case unicode.IsSymbol(r):
			categories["S"]++
		case unicode.IsSpace(r):
			categories["Z"]++
		case unicode.IsControl(r):
			categories["C"]++
		}
		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Println("\ncategory\tcount\n")
	for s, n := range categories {
		fmt.Printf("%s\t\t%d\n", s, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}