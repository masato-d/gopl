package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func wordfreq(f *os.File) map[string]int {
	words := make(map[string]int)
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		words[sc.Text()]++
	}
	return words
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please specify the filename")
	}
	fp, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	fmt.Println(wordfreq(fp))
}