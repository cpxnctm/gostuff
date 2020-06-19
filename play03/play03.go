package main

import "fmt"

func main() {

	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
func foo() func() int {
	x := 2
	y := 3
	return func() int {
		x *= x
		y++
		return x
	}

}
