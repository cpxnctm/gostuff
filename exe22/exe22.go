package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	//use slicing and append to leave: 42, 43, 44, 48, 49, 50, 51 (remove 45-47)
	y := append(x[0:3], x[6:10]...) //this can alsso be done with 	y := append(x[:3], x[6:]...)
	fmt.Println(y)

}
