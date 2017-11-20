package training4_6

import "testing"

func TestCompress(t *testing.T) {
	expected := "abc d e fg h i jk l mn"
	actual := string(compress([]byte("abc d e  fg h    i jk l   mn")))
	if expected != actual {
		t.Errorf("expected:%s actual:%s", expected, actual)
	}
}