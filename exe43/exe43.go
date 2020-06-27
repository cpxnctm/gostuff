package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var wg2 sync.WaitGroup
var mu sync.Mutex

func main() {

	lim := 100
	lim2 := 100
	counter := 0
	counter1 := 0
	wg.Add(lim)
	wg2.Add(lim)

	go func() {
		for i := 0; i < lim; i++ {
			mu.Lock()
			i := counter
			i++
			counter = i
			fmt.Println(counter)
			mu.Unlock()
			wg.Done()
		}
	}()
	go func() {
		for i := 0; i < lim2; i++ {
			mu.Lock()
			i := counter1
			i++
			counter1 = i
			fmt.Println(counter1)
			mu.Unlock()
			wg2.Done()
		}
	}()
	wg.Wait()
	wg2.Wait()
	fmt.Println("Program Exiting")

}
