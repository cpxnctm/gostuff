package main

import (
	"fmt"
)

func main() {

	x := 5
	y := 10
	z := 15
	switch {
	case (x > z):
		fmt.Println("x is greater than z")
	case (x+y == z):
		fmt.Println("x+y=z")

	}
}
