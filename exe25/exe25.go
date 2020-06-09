package main

import (
	"fmt"
)

func main() {
	//create a mapping between an object and corresponding information such as their favorite things. print out the values as well as their position within the slice

	x := map[string][]string{
		`cats`: []string{`meow`, `treats`, `lasers`, `sleep`},
		`dogs`: []string{`woof`, `bones`, `fetching`, `pets`},
		`fish`: []string{`swimming`, `breathing`, `watching`, `floating`},
	}
	for k, v := range x {
		fmt.Printf("For record: %v\n", k)
		for l, v2 := range v {
			fmt.Printf("The entry is: %v and the index position is: %v\n", v2, l)
		}

	}
}
