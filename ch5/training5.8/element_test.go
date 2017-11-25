package training5_8

import (
	"testing"
	"strings"

	"golang.org/x/net/html"
)

func TestElementById(t *testing.T) {
	dom := `<html><head><meta charset="utf-8" /><title>test</title><script type="text/javascript" src="test.js"></script></head><body><h1>test</h1><div><p id="test"><a href="test.html">test</a></p><img src="test.jpg" /></div></body></html>`
	n, _ := html.Parse(strings.NewReader(dom))
	r := ElementById(n, "test")

	if r == nil {
		t.Errorf("result is nil")
		return
	}

	if r.Data != "p" {
		t.Errorf("result:%v", r.Data)
	}
}