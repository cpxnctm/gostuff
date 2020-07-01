package main

import "fmt"

func main() {

	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)
	receive(eve, odd, quit)
	fmt.Println("About to exit")
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case a := <-e:
			fmt.Println("evens:", a)
		case a := <-o:
			fmt.Println("odds:", a)
		case a := <-q:
			fmt.Println("quit:", a)
			return
		}

	}

}
func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0

}
