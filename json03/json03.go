package main

import (
	"encoding/json"
	"fmt"
)

type coffee struct {
	Name   string
	Flavor bool
	Source string
	Buy    bool
	Price  int
}

func main() {
	c1 := coffee{
		Name:   "diner",
		Flavor: false,
		Source: "s.america",
		Buy:    false,
		Price:  5,
	}
	c2 := coffee{
		Name:   "slow roast",
		Flavor: true,
		Source: "s. america",
		Buy:    true,
		Price:  20,
	}
	c3 := coffee{
		Name:   "instant",
		Flavor: false,
		Source: "who knows",
		Buy:    false,
		Price:  5,
	}

	coffees := []coffee{c1, c2, c3}

	xc, err := json.Marshal(coffees)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(string(xc))
	}
}
