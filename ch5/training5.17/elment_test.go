package training5_17

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestElementsByTagNameImg(t *testing.T) {
	dom := `<html><head><title></title></head><body><h1><img src="test.jpg"></h1><h2>test-h2</h2><h3>test-h3</h3><h4>test-h4</h4></body></html>`
	r := strings.NewReader(dom)
	doc, err := html.Parse(r)
	if err != nil {
		t.Error("failed to prepare for testing")
		return
	}
	images := ElementsByTagName(doc, "img")
	if len(images) != 1 {
		t.Errorf("failed to get img element:%+v", images)
		return
	}
	for _, image := range images {
		if image.Data != "img" {
			t.Errorf("failed to get img element:%+v", images)
		}
	}
}

func TestElementsByTagNameHeadings(t *testing.T) {
	dom := `<html><head><title></title></head><body><h1><img src="test.jpg"></h1><h2>test-h2</h2><h3>test-h3</h3><h4>test-h4</h4></body></html>`
	r := strings.NewReader(dom)
	doc, err := html.Parse(r)
	if err != nil {
		t.Error("failed to prepare for testing")
		return
	}
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	if len(headings) != 4 {
		t.Errorf("failed to get img element:%+v", headings)
		return
	}
	for _, heading := range headings {
		if heading.Data != "h1" && heading.Data != "h2" && heading.Data != "h3" && heading.Data != "h4" {
			t.Errorf("failed to get img element:%+v", headings)
		}
	}
}
