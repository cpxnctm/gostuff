package main

import "fmt"

func main() {

	a := make(chan int)

	go send(a)
	receive(a)
	fmt.Println("Exiting Shortly...")
}

func send(b chan<- int) {
	for i := 0; i < 100; i++ {
		b <- i
	}
	close(b)
}
func receive(c <-chan int) {
	for d := range c {
		fmt.Println(d)
		fmt.Printf("%T\t%U\n", d, d)
		fmt.Println("---")

	}
}
