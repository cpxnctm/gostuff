package main

import "fmt"

func main() {

	even := make(chan int)
	odds := make(chan int)
	quit := make(chan int)

	go send(even, odds, quit)
	receive(even, odds, quit)
	fmt.Println("Program Complete")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 150; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}
func receive(e, o, q <-chan int) {
	for {
		select {
		case n := <-e:
			fmt.Println("Even:", n)
		case n := <-o:
			fmt.Println("Odd:", n)
		case n := <-q:
			fmt.Println("Exiting:", n)
			return
		}
	}
}
