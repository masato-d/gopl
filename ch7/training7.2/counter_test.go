package training7_2

import (
	"testing"
	"bytes"
)

func TestCountingWriter(t *testing.T) {
	ts := "test"
	w, i := CountingWriter(bytes.NewBufferString(ts))

	if _, err := w.Write([]byte(ts)); err != nil {
		t.Error("test failed", err)
	}
	if *i != int64(len(ts)) {
		t.Errorf("test failed:%#v", w)
	}

	if _, err := w.Write([]byte(ts)); err != nil {
		t.Error("test failed", err)
	}
	if *i != int64(len(ts))*2 {
		t.Errorf("test failed:%#v", w)
	}
}