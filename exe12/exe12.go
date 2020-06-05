package main

import (
	"fmt"
)

func main() {
	//Print every rune code point of the uppercase alphabet three times. (65-90)

	for i := 65; i <= 90; i++ {
		fmt.Printf("%d\t\n%#U\n%#U\n%#U\n", i, i, i, i)

	}
}
