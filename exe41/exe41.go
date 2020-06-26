package main

import "fmt"

// “The method set of a type determines the INTERFACES that the type implements.....” ~ golang spec
//The method 'saysomething' which takes in the type 'person' is called in the interface 'speak'
// even though the speak method takes in a pointer of person, in func.main, you can still pass in a non-pointer value by calling person.speak
//
//This explanation didn't seem as obvious so I dug a bit more and found this answer on stackoverflow which makes more sense (at first sight). what do you thing?
//https://stackoverflow.com/questions/42477951/what-is-the-method-set-of-sync-waitgroup
//This is in Spec: Calls:

//If x is addressable and &x's method set contains m, x.m() is shorthand for (&x).m().
//So when you have a value that is addressable (a variable is addressable), you may call any methods that have pointer receiver on non-pointer values, and the compiler will automatically take the address and use that as the receiver value.

type person struct {
	name  string
	age   int
	hobby string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("I am:", p.name, "I am:", p.age, "I like:", p.hobby)
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
