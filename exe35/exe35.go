package main

import (
	"fmt"
)

type square struct {
	length, width float64
}

func (sq square) area() float64 {
	return sq.length * sq.width
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}
func main() {
	x := square{
		length: 123,
		width:  456,
	}
	info(x)
	fmt.Printf("%b\n", x)
}
