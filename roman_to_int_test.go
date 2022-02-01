package main

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	v := "XI"
	res := romanToInt(v)
	if res != 11 {
		t.Errorf("%d as the result", res)
	}
}
