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
	fmt.Println(&x)
	bar(&x)
	fmt.Println(x)
}
func foo(p *person) {
	p.name = "Bootsy Kitty"

}
func bar(p2 *person) {
	p2.name = "Buff Kitty"
}
