package main

import (
	"fmt"
)

//Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4, the result is %v\n", i, i%4) // don't forget that you don't need to write an entire function to get the result. here, we do the math right in the print function
	}
}
