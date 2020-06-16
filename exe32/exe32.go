package main

import "fmt"

func main() {
	x1 := foo([]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}...)
	ii := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	x2 := bar(ii)
	fmt.Println(x1)
	fmt.Println(x2)
}

func foo(x ...int) int {
	a := 0
	for _, v := range x {
		a += v
	}
	return a
}

func bar(y []int) int {
	b := 0
	for _, v := range y {
		b += v
	}
	return b
}
