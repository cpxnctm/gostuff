package main

import "fmt"

type person struct {
	name string
}

func main() {
	x := person{
		name: "Beta Pupper",
	}
	fmt.Println(x)
	foo(&x)
	fmt.Println(x)
}
func foo(p *person) {
	p.name = "Bootsy Kitty"

}
