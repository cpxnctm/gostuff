package main

import "fmt"

func main() {

	x := []int{2, 4, 6, 8, 10, 12, 14}
	fmt.Println(x[1:4])
	// print the index values and the integer values
	y := []int{100, 200, 300, 400, 500}
	x = append(x, y...)
	fmt.Println(x)
	for a, b := range x {
		fmt.Printf("At index location %v is: %v\n", a, b)

		for i := 0; i < 7; i++ {
			fmt.Printf("At index location %v is: %v\n", i, x[i])
		}
	}

}
