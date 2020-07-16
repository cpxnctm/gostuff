package mymath

import (
	"fmt"
	"testing"
)


func TestCenteredAvg(t *testing.T) {
	type test struct{
		data []int
		answer float64
	}
	tst := []test{
		test{[]int{1,2,3}, 2},
		test{[]int{4,5,6}, 5},
		test{[]int{7,8,9}, 8},
	}
	for _,v := range tst{
		d := v.data
		a := v.answer
		r := CenteredAvg(d)
		if r != a{
	t.Errorf("Expected: %v, got: %v", a,r )
		}
	}
}

func ExampleCenteredAvg() {
	a := []int{10,20,30,40}
	fmt.Println(CenteredAvg(a))
	//Output:
	//25
}

func BenchmarkCenteredAvg(b *testing.B) {
	bn := []int{10,20,30,40}
	for i :=0; i < b.N; i++{
		CenteredAvg(bn)
	}
}