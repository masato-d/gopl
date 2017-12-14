package training7_4

import (
	"io"
	"strings"
)

type HtmlReader string

func (r *HtmlReader) Read(p []byte) (n int, err error) {
	n, err = strings.NewReader(string(*r)).Read(p)
	return
}

func NewReeader(s string) io.Reader {
	r := HtmlReader(s)
	return &r
}