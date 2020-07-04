package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func receive(n <-chan int) {
	for i := range n {
		fmt.Println(i)
	}
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c) //don't forget to close the channel
	}()
	return c // don't forget to return the channel
}
