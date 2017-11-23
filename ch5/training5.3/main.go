package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func showText(n *html.Node) {
	if n == nil {
		return
	}

	if n.Data == "style" || n.Data == "script" {
		return
	}

	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}

	showText(n.FirstChild)
	showText(n.NextSibling)
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "showText:%v\n", err)
		os.Exit(1)
	}

	showText(doc)
}