package main

import (
	"fmt"
)

func main() {

	//Using a COMPOSITE LITERAL: create a SLICE of TYPE int and assign 10 values

	x := []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}
	for r := range x {
		fmt.Printf("%d\n", x[r]) // range over the slice and print out the values
		fmt.Printf("%T\n", x[r]) //range over the slice and print out the type
	}

}
