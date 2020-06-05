package main

import (
	"fmt"
)

func main() {
	//Print every rune code point of the uppercase alphabet three times. (65-90)

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for r := 0; r < 3; r++ {
			fmt.Printf("\t%#U\n", i)
		}

	}
}
