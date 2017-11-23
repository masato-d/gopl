package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var domCount = make(map[string]int)

func count(n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		domCount[n.Data]++
	}

	count(n.FirstChild)
	count(n.NextSibling)
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	count(doc)
	for dom, num := range domCount {
		fmt.Printf("%s: %d\n", dom, num)
	}
}