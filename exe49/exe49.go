package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c)

	fmt.Println("about to exit")
}

//use a for loop statement to pull the information off of the channel

func gen(q <-chan int) <-chan int {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
	return c
}

func receive(n <-chan int) {
	for i := range n {
		fmt.Println(i)
	}
}
