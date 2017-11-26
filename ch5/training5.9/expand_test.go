package training5_9

import "testing"

func TestExpand(t *testing.T) {
	s := "test $foo testtest $hoge"
	expected := "test foo-expanded testtest hoge-expanded"
	actual := expand(s, func(str string) string {
		return str + "-expanded"
	})
	if expected != actual {
		t.Errorf("expected:%s actual:%s", expected, actual)
	}
}