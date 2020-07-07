package main

import "fmt"

// assign a func to a variable and then call the func

func main() {
	a := func(x int) {
		fmt.Println(x)
	}
	a(10)
}
