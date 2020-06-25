package main

import (
	"fmt"
	"runtime"
	"sync"
)

// launch two additional goroutines
//each goroutine should print something out

var counter int
var counter2 int
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	fmt.Println("Number of Cores: ", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	lim := 50
	lim2 := 50
	wg.Add(lim)

	for a := 0; a < lim; a++ {
		go func() {
			mu.Lock()
			b := counter
			fmt.Printf("The limit is: %v\n", lim)
			counter = b
			mu.Unlock()
			fmt.Println("-----------------------")
			wg.Done()
		}()
	}
	fmt.Println("Number of Cores: ", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("-----------------------")

	for c := 0; c < lim2; c++ {
		go func() {
			mu.Lock()
			d := counter2
			fmt.Println("This is the second Goroutine:", runtime.NumGoroutine())
			counter2 = d
			mu.Unlock()
			fmt.Println("-----------------------")
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Routines Complete")

}
