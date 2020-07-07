package main

import "fmt"

func main() {
	//Create a program that shows the “if statement” in action.

	const x = 20
	const y = 10

	if x > y {
		fmt.Println("something isn't right here...")
	} else if x < y {
		fmt.Println("looks we did alright...")
	} else if x == y {
		fmt.Println("we really did something wrong...")
	} else {
		fmt.Println("numbers are hard")
	}
}
