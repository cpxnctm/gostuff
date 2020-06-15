package main

import (
	"fmt"
)

func main() {
	t := evenSum(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...) // t is equal to the evenSum func which calls the sum func and passess the slice of int as an arg
	fmt.Println(t)                                         //prints t
}

func sum(x ...int) int { //create the sum func
	n := 0
	for _, v := range x { //toss the index, but keep the value and for range x, add up the values
		n += v
	}
	return n //return the value n
}
func evenSum(f func(x ...int) int, y ...int) int { // the evenSum func has parameters of unlimited ints assigned to f and returns an int. it also takes in unlimited ints assigned to y
	var xi []int          //creates a slice of int assigned to var xi
	for _, v := range y { //range over the values passed in from y and tosses out the index
		if v%2 == 0 { //to find the even numbers, divide by two and check for a zero remainder
			xi = append(xi, v) // if the remainder is zero, then append the value to the slice: xi
		}
	}
	total := f(xi...) // calls the function f and passes in the int values from xi
	return total      //returns the total to evenSum
}
