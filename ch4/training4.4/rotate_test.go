package training4_4

import (
	"testing"
	"fmt"
)

func TestRotate(t *testing.T) {
	expected := []int{2, 3, 4, 5, 0, 1}
	actual := rotate([]int{0, 1, 2, 3, 4, 5}, 2)
	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", actual) {
		t.Errorf("expected:%s actual:%s", expected, actual)
	}
}