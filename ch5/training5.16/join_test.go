package training5_16

import (
	"testing"
)

func TestCustomJoin(t *testing.T) {
	expected := "I am in a test function"
	actual := customJoin(" ", "I", "am", "in", "a", "test", "function")
	if expected != actual {
		t.Errorf("expected:%s actual:%s", expected, actual)
	}
}