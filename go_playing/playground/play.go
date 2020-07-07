package main

import "fmt"

func main() {
	x := []int{20, 40, 60, 80, 100, 120}
	fmt.Println(x)
	fmt.Println(x[2])
	fmt.Println(x[0:3])
	for i := 0; i <= 5; i++ {
		fmt.Println(i, x[i])

	}

}
