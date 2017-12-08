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

func TestRemove(t *testing.T) {
	var x IntSet

	x.Add(1)
	x.Add(2)
	if !x.Has(2) {
		t.Error("x should have 2")
	}

	x.Remove(2)
	if x.Has(2) {
		t.Error("x should not have 2")
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

func TestCopy(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(2)

	y := x.Copy()

	if x.Len() != y.Len() || !y.Has(1) || !y.Has(2) {
		t.Errorf("failed to copy:%v", y)
	}

	y.Add(3)
	if x.Has(3) {
		t.Errorf("failed to copy:%v\n%v", x, y)
	}
}