package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, val := range counts {
		n := val["count"]
		var files, sep string
		if n > 1 {
			for k, _ := range val {
				if k != "count" {
					files += sep + k
					sep = " "
				}
			}
			fmt.Printf("%d\t%s\t%s\n", n, line, files)
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if _, ok := counts[input.Text()]; ok {
			counts[input.Text()]["count"]++
		} else {
			counts[input.Text()] = make(map[string]int)
			counts[input.Text()]["count"] = 1
		}
		counts[input.Text()][f.Name()]++
	}
}

