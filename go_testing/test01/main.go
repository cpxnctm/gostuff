package main

import "fmt"

func main() {
	fmt.Println(mult(2, 3, 4))
	fmt.Println(mult(5, 6, 7))
	fmt.Println(mult(8, 9, 10))

}

func mult(n ...int) int {
	y := 10
	for _, v := range n {
		y *= v
	}
	return y
}
