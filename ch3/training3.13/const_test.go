package training3_13

import "testing"

func TestKB(t *testing.T) {
	if KB != 1000 {
		t.Errorf("KB is %d", KB)
	}
}

func TestMB(t *testing.T) {
	if MB != 1000000 {
		t.Errorf("MB is %d", MB)
	}
}

func TestGB(t *testing.T) {
	if GB != 1000000000 {
		t.Errorf("GB is %d", GB)
	}
}

func TestTB(t *testing.T) {
	if TB != 1000000000000 {
		t.Errorf("TB is %d", TB)
	}
}

func TestPB(t *testing.T) {
	if PB != 1000000000000000 {
		t.Errorf("PB is %d", PB)
	}
}

func TestEB(t *testing.T) {
	if EB != 1000000000000000000 {
		t.Errorf("EB is %d", EB)
	}
}