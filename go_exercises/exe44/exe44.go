package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {

	lim := 50
	var counter int64
	wg.Add(lim)

	go func() {
		for i := 0; i < lim; i++ {
			atomic.AddInt64(&counter, 1)
			i := counter
			i++
			counter = i
			fmt.Println(counter)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}
	}()
	wg.Wait()
	fmt.Println("Program Exiting")

}
