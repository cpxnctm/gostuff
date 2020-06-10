package main

import (
	"fmt"
)

func main() {

	// create a slice with TYPE int and use SLICE and APPEND to delete values

	x := []int{5, 10, 15, 20, 25, 30, 35, 40}
	fmt.Println(x)
	y := append(x[0:2], x[6:]...)
	fmt.Println(y)
}
