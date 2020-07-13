package main

import "testing"

func TestMult(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}
	tst := []test{
		test{[]int{2, 3, 4}, 240},
		test{[]int{5, 6, 7}, 2100},
		test{[]int{8, 9, 10}, 7200},
	}
	for _, v := range tst {
		x := mult(v.data...)
		if x != v.answer {
			t.Error("Expected:", v.answer, "Got:", x)
		}

	}
}
