package main

import "fmt"

func main() {
	a := []string{"dog", "cat", "mouse"}
	b := []string{"car", "bike", "scooter"}
	fmt.Println(a)
	fmt.Println(b)
	xab := [][]string{a, b}
	fmt.Println(xab)
	c := (b[0:1])
	fmt.Println(c)
}
