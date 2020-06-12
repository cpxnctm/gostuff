package main

import (
	"fmt"
)

func main() {
	// create a type named person that is a struct and store first, last and favorite ice cream flavor

	type person struct {
		first      string
		last       string
		favFlavors string
	}
	x := person{
		first:      `bob`,
		last:       `jones`,
		favFlavors: `mint`,
	}
	fmt.Println(x)
}
