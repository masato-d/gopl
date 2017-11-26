package training5_17

import "golang.org/x/net/html"

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var elements []*html.Node
	f := func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, tag := range name {
				if n.Data == tag {
					elements = append(elements, n)
				}
			}
		}
	}
	forEachNode(doc, f, nil)
	return elements
}

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
