package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var counter int

func main() {

	const lim = 100
	fmt.Println("Starting Routines")
	wg.Add(lim)

	for i := 0; i < lim; i++ {
		go func() {
			a := counter
			runtime.Gosched()
			a++
			counter = a
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
		fmt.Println("Cores:", runtime.NumCPU())
	}

	wg.Wait()
	fmt.Println("Goroutines completed")
}
