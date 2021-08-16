package gomodmath

import "testing"

func TestAdd(t *testing.T) {
	val := Add(1, 2)
	if val != 3 {
		t.Fatal("add func failed")
	}
}

func TestSub(t *testing.T) {
	val := Sub(3, 2)
	if val != 1 {
		t.Fatal("sub func failed")
	}
}
