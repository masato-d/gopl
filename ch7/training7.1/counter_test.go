package training7_1

import (
	"fmt"
	"testing"
)

func TestWordCounter_Write(t *testing.T) {
	var c WordCounter
	fmt.Fprint(&c, "hoge fuga foo bar")

	if fmt.Sprintf("%d", c) != "4" {
		t.Errorf("failed to write: %#v", c)
	}
}

func TestLineCounter_Write(t *testing.T) {
	var c LineCounter
	fmt.Fprint(&c, "hoge\nfuga\nfoo\nbar")

	if fmt.Sprintf("%d", c) != "4" {
		t.Errorf("failed to write: %#v", c)
	}
}