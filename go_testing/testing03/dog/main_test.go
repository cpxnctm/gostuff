package dog

import (
	"testing"
)

func TestYearsTwo(t *testing.T) {
	v := YearsTwo(7)
	if v != 49 {
		t.Errorf("Expected: %v, Got: %v", v, t)
	}

}

func BenchTestYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(7)
	}
}
