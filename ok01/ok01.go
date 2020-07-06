package main

import "fmt"

func main() {
	c := make(chan int)
	a := make(chan int)

	go func() {
		c <- 10
		close(c)
	}()
	v, ok := <-c
	fmt.Println(v, ok)
	v, ok = <-c
	fmt.Println(v, ok)

	go func() {
		a <- 5
		close(a)
	}()
	v, ok = <-a
	fmt.Println(v, ok)
	v, ok = <-a
	fmt.Println(v, ok)
}
