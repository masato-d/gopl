package training6_1

import (
	"testing"
	"fmt"
)

func TestLen(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)

	expected := 3
	if x.Len() != expected {
		t.Errorf("expected:%d actual:%d", expected, x.Len())
	}
}

func TestClear(t *testing.T) {
	var x IntSet

	x.Add(1)
	x.Add(2)
	x.Clear()
	if fmt.Sprintf("%v", x.String()) != "{}" {
		t.Errorf("failed to clear: %v", x)
	}
}