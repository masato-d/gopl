package training5_8

import (
	"golang.org/x/net/html"
)

var node *html.Node
var searchID string

func ElementById(doc *html.Node, id string) *html.Node {
	searchID = id
	forEachNode(doc, shouldContinue, nil)
	return node
}

func forEachNode(n *html.Node, pre, post func() bool ) {
	if pre != nil {
		if ok := pre(); !ok {
			return
		}
	}

	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == searchID {
				node = n
				return
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post()
	}
}

func shouldContinue() bool {
	if node != nil {
		return false
	}
	return true
}