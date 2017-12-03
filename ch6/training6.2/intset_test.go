package training6_2

import "testing"

func TestAddAll(t *testing.T) {
	var s IntSet
	s.AddAll(1, 2, 3)
	if !s.Has(1) || !s.Has(2) || !s.Has(3) {
		t.Errorf("failed to test AddAll: %v", s)
	}
}