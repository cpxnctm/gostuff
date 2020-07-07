package main

import "fmt"

type animal struct {
	first string
	last  string
	age   int
}

func (s animal) speak() { // scope is important...
	fmt.Printf("My name is %s, my last name is, %s and I'm %v years old\n", s.first, s.last, s.age)
}
func main() {

	x := animal{
		first: "bootsy",
		last:  "kitty",
		age:   9,
	}
	y := animal{
		first: "beta",
		last:  "doggo",
		age:   1,
	}
	x.speak() //don't forget to add the parens in order to execute the method
	y.speak()
}
