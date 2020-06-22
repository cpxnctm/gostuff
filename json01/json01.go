package main

import (
	"encoding/json"
	"fmt"
)

type pet struct {
	Name   string
	Animal string
	Age    int
	Cute   bool
}

func main() {
	p1 := pet{
		Name:   "Bootsy",
		Animal: "Kitty",
		Age:    9,
		Cute:   true,
	}
	p2 := pet{
		Name:   "Beta",
		Animal: "Pupper",
		Age:    1,
		Cute:   true,
	}
	pets := []pet{p1, p2}
	fmt.Println(pets)

	a, err := json.Marshal(pets)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(a))

}
