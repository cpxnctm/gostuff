package main

import "fmt"

type pet struct {
	name string
}

func main() {
	x := pet{
		name: "Beta Pupper",
	}
	fmt.Println(x)
	foo(&x)
	fmt.Println(&x)
	bar(&x)
	fmt.Println(x)
}
func foo(p *pet) {
	p.name = "Bootsy Kitty"

}
func bar(p2 *pet) {
	p2.name = "Buff Kitty"
}
