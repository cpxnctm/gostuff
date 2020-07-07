package main

import (
	"fmt"
)

func main() {
	//anonymous struct

	s := struct {
		name  string
		food  bool
		value int
	}{
		name:  `carrot`,
		food:  true,
		value: 20,
	}
	fmt.Println(s)
}
