package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() { //create an anonymous goroutine to send the value onto the channel
		c <- 42
	}()
	fmt.Println(<-c)
}
