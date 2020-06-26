package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {

	lim := 100
	counter := 0
	wg.Add(lim)

	go func() {
		for i := 0; i < lim; i++ {
			mu.Lock()
			i = counter
			i++
			counter = i
			fmt.Println(counter)
			mu.Unlock()
			wg.Done()
		}
		fmt.Println("The current value is:", counter)
	}()

	wg.Wait()
	fmt.Println("Program Exiting")
}
