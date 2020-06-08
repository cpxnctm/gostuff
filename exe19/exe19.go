package main

import (
	"fmt"
)

func main() {

	var x [5]int
	fmt.Println((x))
	x[0] = 20
	x[1] = 25
	x[2] = 30
	x[3] = 35
	x[4] = 40
	fmt.Println(x)
	//fmt.Printf("%T%v\n%T%v\n%T%v\n%T%v\n%T%v\t", x[0], x[0], x[1], x[1], x[2], x[2], x[3], x[3], x[4], x[4])

	for r := range x {
		fmt.Printf("%T%v\n", x[r], x[r])
	}

}
