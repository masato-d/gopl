package training4_5

import (
	"testing"
	"fmt"
)

func TestFilter(t *testing.T) {
	expected := []string{"a", "b", "c", "a", "c", "d", "e"}
	actual := filter([]string{"a", "b", "b", "c", "a", "c", "d", "d", "d", "e"})
	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", actual) {
		t.Errorf("expected:%v actual:%v", expected, actual)
	}
}