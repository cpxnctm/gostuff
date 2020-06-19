package main

import (
	"fmt"
)

func main() {

	x := 25
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(*&x)
	foo(&x)
}

func foo(n *int) {
	y := n
	fmt.Println("break")
	*y = 100
	fmt.Println(&y)
	fmt.Println(*y)
}
