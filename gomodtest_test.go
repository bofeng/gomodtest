package gomodmath

import "testing"

func TestAdd(t *testing.T) {
	val := Add(1, 2)
	if val != 3 {
		t.Fatal("add func failed")
	}
}
