package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("specify URL")
	}

	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("cannot access %s: %s", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	forEachNode(doc, startElement, endElement)
}