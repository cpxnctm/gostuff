package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var counter int
var mu sync.Mutex

func main() {

	const lim = 100
	fmt.Println("Starting Routines")
	wg.Add(lim)

	for i := 0; i < lim; i++ {
		go func() {
			mu.Lock()
			a := counter
			runtime.Gosched()
			a++
			counter = a
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
		fmt.Println("Cores:", runtime.NumCPU())
	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println("Goroutines completed")
}
