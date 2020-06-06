package main

import "fmt"

func main() {
	x := []int{20, 40, 60, 80, 100, 120} //define the slice
	fmt.Println(x)
	fmt.Println(x[2])         //print the value at index 2 (remember 0 base)
	fmt.Println(x[0:3])       //print the values from index 0 through 3 (20,40,60)
	for i := 0; i <= 5; i++ { //loop through the slice and print the index number and the value at said index
		fmt.Println(i, x[i])

		for a, b := range x { // loop through the range specified in 'x' and print the index and value at index
			fmt.Println(a, b)
		}

	}

}
