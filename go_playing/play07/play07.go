package main

import (
	"fmt"
)

func main() {

	a := 100
	fmt.Println(a)
	fmt.Println(&a)
	foo(&a)
	a = 120
	fmt.Println(a)
	fmt.Println(&a)
	a = 5
	fmt.Println(a)
	fmt.Println(&a)

}
func foo(p *int) {
	fmt.Println(p)

}
