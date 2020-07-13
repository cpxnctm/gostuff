package main

import "testing"

func TestMult(t *testing.T) {
	x := mult(2, 3, 4)
	y := mult(5, 6, 7)
	z := mult(8, 9, 10)
	if x != 240 {
		t.Errorf("Expected 240, got %v", x)
	}
	if y != 2100 {
		t.Errorf("Expected 2100, got %v", y)
	}
	if z != 7200 {
		t.Errorf("Expected 7200, got %v", z)
	}
}
