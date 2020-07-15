package say

import (
	"fmt"
	"testing"
)

func TestKind(t *testing.T) {
	a := Kind("Bootsy")
	if a != "Have a great day, Bootsy" {
		t.Error("Expected", t, "Got", a)
	}
}

func ExampleKind() {
	fmt.Println(Kind("Bootsy"))
	//Output:
	//Have a great day, Bootsy
}

func BenchKind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Kind("Bootsy")

	}
}
