package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var wg2 sync.WaitGroup
var wg3 sync.WaitGroup

func main() {
	var lim = 50
	var a int

	wg.Add(lim)
	wg2.Add(lim)
	wg2.Add(lim)

	go func() {
		i := a
		for i := 0; i < lim; i++ {
			a = i
			wg.Done()
		}
		fmt.Println(i)

	}()
	fmt.Println("Gorountines:\t", runtime.NumGoroutine())
	go func() {
		b := a
		for b := 0; b < lim; b++ {
			a = b
			wg2.Done()
		}
		fmt.Println(b)

	}()
	go func() {
		c := a
		for c := 0; c < lim; c++ {
			a = c
			wg3.Done()
		}
		fmt.Println(c)

	}()
	fmt.Println("Process Complete")
	wg.Wait()
	wg2.Wait()
	wg3.Wait()
}
