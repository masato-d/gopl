package training5_19

import "testing"

func TestNoReturn(t *testing.T) {
	expectd := "not zero value"
	actual := noReturn()
	if expectd != actual {
		t.Errorf("expected:%s actual:%s", expectd, actual)
	}
}