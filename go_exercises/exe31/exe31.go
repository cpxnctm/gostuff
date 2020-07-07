package main

import "fmt"

//create a function named foo that returns an int
//create a function named bar that returns an int and a string
// call both functions and print their results
func main() {
	x := foo()
	fmt.Println(x)
	y, z := bar()
	fmt.Println(y, z)
}

func foo() int {
	a := 0
	for a := 0; a <= 50; a++ {
	}
	return a
}

func bar() (int, string) {
	b := ("The world is a big place!")
	c := 5
	return c, b
}
