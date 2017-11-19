package training3_12

import "testing"

func TestAnagramTrue(t *testing.T) {
	w1 := "stop"
	w2 := "post"
	res := anagram(w1, w2)
	if res != true {
		t.Errorf("the two words is NOT an anagram: %s, %s", w1, w2)
	}
}

func TestAnagramFalse(t *testing.T) {
	w1 := "shop"
	w2 := "post"
	res := anagram(w1, w2)
	if res != false {
		t.Errorf("the two words is an anagram: %s, %s", w1, w2)
	}
}