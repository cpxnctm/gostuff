package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println("how are you doing today?")
	bootsy()
	beta()

}
func bootsy() {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
		}
		fmt.Println(i)
	}
}
func beta() {
	fmt.Println("today is going woof!")
}
