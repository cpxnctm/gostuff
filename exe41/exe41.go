package main

import "fmt"

// “The method set of a type determines the INTERFACES that the type implements.....” ~ golang spec
//The method 'walk' which takes in the type 'person' is called in the interface 'speak'
// even though the speak method takes in a pointer of person, in func.main, you can still pass in a non-pointer value by calling person.speak

type person struct {
	name  string
	age   int
	hobby string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("I am:", p)
}

func saySomething(h human) {
	h.speak()
}

func main() {

	a := person{
		name:  "Beta",
		age:   1,
		hobby: "ice cubes",
	}
	saySomething(&a)
	a.speak()
}
