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
func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}

}
func TestUseCount(t *testing.T) {
	a := UseCount("cat dog fox fox fox cat dog dog")
	for k, v := range a{
		switch k {
		case "cat":
			if v != 2{
				t.Error("Expected 2, got:", v)
			}
		case "dog":
			if v != 3{
				t.Error("Expected 3, got:", v)

			}
		case "fox":
			if v != 3{
				t.Error("Expected 3, got:", v)
			}
		}
	}

}


