package gomodtest

import "testing"

func TestAdd(t *testing.T) {
	val := add(1, 2)
	if val != 3 {
		t.Fatal("add func failed")
	}
}
