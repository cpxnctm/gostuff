package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rectangle struct {
	length float64
	width  float64
}
type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius

}
func (r rectangle) area() float64 {
	return r.length * r.width
}
func info(sh shape) {
	fmt.Println(sh.area())
}

func main() {
	a := circle{
		radius: 5.4321,
	}
	b := rectangle{
		length: 2.468,
		width:  5.4321,
	}
	c := rectangle{
		length: 100.50,
		width:  25.234,
	}
	info(a)
	info(b)
	info(c)

}
