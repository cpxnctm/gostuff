package main

import (
	"fmt"
)

type person struct {
	first string
	last string
}
type secretAgent struct {
	person
	ltk bool
}

 func (s, secretAgent) speak () {
	fmt.Println("Hello, ", s.first, s.last)
 }
func main (){
 sa1 := secretAgent {
	 person:person {
		 "james",
		 "bond",
	 }
	 ltk: true
	}
	fmt.Println(sa1)
	sa1.speak

	
	
}