package training4_3

import "testing"

func TestReverse(t *testing.T) {
	expected := [6]int{5, 4, 3, 2, 1, 0}
	a := &[6]int{0, 1, 2, 3, 4, 5}
	reverse(a)
	if expected != *a {
		t.Errorf("expected:%v actual:%v", expected, *a)
	}
}