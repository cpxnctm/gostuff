package main

import "fmt"

func main() {
	func(x int) {
		fmt.Println("Good news everyone!", x)
	}(20)

}
