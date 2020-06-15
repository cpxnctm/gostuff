package main

import (
	"fmt"
)

//create a function named foo that returns an int
//create a function named bar that returns an int and a string
// call both functions and print their results
func main() {
	a := foo(2)
	fmt.Println(a)
	fmt.Println(bar)
}
func foo(x int) int {
	for i := 0; i < 10; i++ {
		return i
	}
}
func bar() (string, int) {
	b := "i'm writing code"
	c := 451
	return "", b
	return "", c
}
