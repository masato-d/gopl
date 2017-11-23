package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := "https://golang.org"

	words, images, err := CountWordsAndImages(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("words:%d images:%d\n", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML:%s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	r1 := strings.NewReader(n.Data)
	r2 := strings.NewReader(n.Namespace)

	var s []byte
	for _, attr := range n.Attr {
		s = append(s, []byte(attr.Key)...)
		s = append(s, []byte(attr.Val)...)
		s = append(s, []byte(attr.Namespace)...)
	}
	r3 := bytes.NewReader(s)

	r := io.MultiReader(r1, r2, r3)

	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		words++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}

	return
}