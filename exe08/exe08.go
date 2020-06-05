package main

import "fmt"

func main() {

	//_ = iota
	a := ((25) * (25252525))
	fmt.Printf("The values of a are: %d\t%b\t%x\n", a, a, a)

	b := (a << 1)
	//fmt.Println(b)
	fmt.Printf("The values of b are: %d\t%b\t%x\n", b, b, b)

}
