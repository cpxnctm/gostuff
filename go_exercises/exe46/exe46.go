package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1) //creates a buffer channel to hold the value so it can be printed out without the need for a goroutine

	c <- 42

	fmt.Println(<-c)
}
