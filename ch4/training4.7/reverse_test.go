package training4_7

import "testing"

func TestReverse(t *testing.T) {
	expected := "esrever"
	s := []byte("reverse")
	reverse(s)
	if expected != string(s) {
		t.Errorf("expected:%s actual:%s", expected, string(s))
	}
}