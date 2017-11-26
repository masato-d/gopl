package training5_15

import "testing"

func TestMax1(t *testing.T) {
	expected := 10
	actual := max1(5, 10, 2, 8, 9, 3, 7)
	if expected != actual {
		t.Errorf("expected:%d actual:%d", expected, actual)
	}
}

func TestMax1Empty(t *testing.T) {
	expected := 0
	actual := max1()
	if expected != actual {
		t.Errorf("expected:%d actual:%d", expected, actual)
	}
}

func TestMax2(t *testing.T) {
	expected := 10
	actual := max2(5, 10, 2, 8, 9, 3, 7)
	if expected != actual {
		t.Errorf("expected:%d actual:%d", expected, actual)
	}
}

func TestMax2Empty(t *testing.T) {
	expected := 10
	actual := max2(10)
	if expected != actual {
		t.Errorf("expected:%d actual:%d", expected, actual)
	}
}