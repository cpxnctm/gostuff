package main

import "fmt"

func main() {
	c := make(chan int)

	for l := 0; l < 10; l++ {
		go func() {
			for n := 0; n < 10; n++ {
				c <- n
			}
		}()
	}
	for v := 0; v < 100; v++ {
		fmt.Println(v, <-c)
	}
}
