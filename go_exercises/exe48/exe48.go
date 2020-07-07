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

func gen() <-chan int { //since it was specified that there would be a return, don't forget to send it back
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c) //don't forget to close the channel
	}()
	return c // don't forget to return the channel
}
