package training3_10

import (
	"testing"
)

func TestComma(t *testing.T) {
	s := "1234567890"
	expected := "1,234,567,890"
	actual := comma(s)
	if expected != actual {
		t.Errorf("expected:%s actual:%s", expected, actual)
	}
}
