package main

import "fmt"

func main() {
	a := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			a <- i
		}
		close(a)
	}()
	for b := range a {
		fmt.Println(b)
	}

}
