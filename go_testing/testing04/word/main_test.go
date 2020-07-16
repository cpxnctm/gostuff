package word

import (
	"fmt"
	"github.com/cpxnctm/gostuff/go_testing/testing04/quote"
	"testing"
)

func ExampleCount(){
	fmt.Println(Count(quote.SunAlso))
	//Output:
	//1348
}

func TestCount(t *testing.T){
	a := Count(quote.SunAlso)
	v := 1348
	if a != v {
		t.Error("Expected 1348, got: ", a)
	}

}
func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++{
		Count(quote.SunAlso)
	}
}
func ExampleBetter(){
	fmt.Println(Better(quote.SunAlso))
	//Output:
	//1348
}
func TestBetter(t *testing.T){
	a := Better(quote.SunAlso)
	v := 1348
	if a != v {
		t.Error("Expected 1348, got: ", a)
	}

}
func BenchmarkBetter(b *testing.B) {
	for i := 0; i < b.N; i++{
		Better(quote.SunAlso)
	}
}




