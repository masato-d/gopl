package training3_11

import "testing"

func TestComma(t *testing.T) {
	expected := "123,456,789.0"
	actual := comma("123456789.0")
	if expected != actual {
		t.Errorf("expected:%s actual:%s", expected, actual)
	}
}