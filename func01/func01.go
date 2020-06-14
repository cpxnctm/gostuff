package main

import (
	"fmt"
)

type thing struct {
	name string
	age  int
}

func (t thing) hey() {
	fmt.Println("Hey there, ", t.name)
}
func main() {
	x := thing{
		name: "bob",
		age:  30,
	}
	fmt.Println(x)
	x.hey()

}
