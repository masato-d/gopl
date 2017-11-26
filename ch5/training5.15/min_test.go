package training5_15

import "testing"

func TestMin1(t *testing.T) {
	expected := -10
	actual := min1(8, 2, 9, 0, -4, 10, -6, -10)
	if expected != actual {
		t.Errorf("expected:%d actual:%d", expected, actual)
	}
}

func TestMin1Empty(t *testing.T) {
	expected := 0
	actual := min1()
	if expected != actual {
		t.Errorf("expected:%d actual:%d", expected, actual)
	}
}

func TestMin2(t *testing.T) {
	expected := -10
	actual := min2(8, 2, 9, 0, -4, 10, -6, -10)
	if expected != actual {
		t.Errorf("expected:%d actual:%d", expected, actual)
	}
}

func TestMin2Empty(t *testing.T) {
	expected := -5
	actual := min2(-5)
	if expected != actual {
		t.Errorf("expected:%d actual:%d", expected, actual)
	}
}