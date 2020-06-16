package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type square struct {
	length float64
}

func (c circle) area() float64 { //functions that pull in the types and return a float64
	return math.Pi * c.radius * c.radius
}
func (s square) area() float64 {
	return s.length * s.length
}

type shape interface { //the shape type that interfaces area
	area() float64
}

func info(s shape) { //prints the area
	fmt.Println(s.area())
}
func main() {
	circ := circle{ //defining the values for the radius and length required for the calculations
		radius: 12,
	}
	sq := square{
		length: 25,
	}
	info(circ) //pass in the args circ and sq into the info function to print
	info(sq)
}
