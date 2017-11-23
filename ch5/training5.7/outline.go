package training5_7

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

var res []string

func outline(dom string) string {
	doc, err := html.Parse(strings.NewReader(dom))
	if err != nil {
		log.Fatal(err)
	}
	forEachNode(doc, startElement, endElement)

	return strings.Join(res, "")
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) string) {
	if pre != nil {
		res = append(res, pre(n))
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		res = append(res, post(n))
	}
}

var depth int

func startElement(n *html.Node) string {
	if n.Type == html.ElementNode {
		var attrs []string
		for _, attr := range n.Attr {
			attrs = append(attrs, fmt.Sprintf("%s=\"%s\"", attr.Key, attr.Val))
		}

		var ct string
		if n.FirstChild == nil && n.Data != "script" {
			ct = "/>"
		} else {
			ct = ">"
		}

		var a string
		if len(attrs) > 0 {
			a = fmt.Sprintf(" %s", strings.Join(attrs, " "))
		}

		s := fmt.Sprintf("%*s<%s%s%s\n", depth*2, "", n.Data, a, ct)
		depth++
		return s
	}
	return ""
}

func endElement(n *html.Node) string {
	if n.Type == html.ElementNode && n.FirstChild == nil && n.Data != "script" {
		depth--
		return ""
	}
	if n.Type == html.ElementNode {
		depth--
		return fmt.Sprintf("%*s</%s>\n", depth*2, "", n.Data)
	}
	return ""
}