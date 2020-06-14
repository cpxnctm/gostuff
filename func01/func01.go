package main

import (
	"fmt"
)

type thing struct {
	name string
	age  int
}

func (t thing) hey() { //method definition
	fmt.Println("Hey there,", t.name)
}
func main() {
	x := thing{
		name: "bob",
		age:  30,
	}
	y := thing{
		name: "joe",
		age:  31,
	}
	fmt.Println(x)
	x.hey() //calling the method. the function that it is attached to is called as func.method
	y.hey()

}
