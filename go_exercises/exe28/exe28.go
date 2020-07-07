package main

import (
	"fmt"
)

func main() {

	person := map[string][]string{
		`first`:     []string{`bob`},
		`last`:      []string{`jones`},
		`favFlavor`: []string{`mint`, `vanilla`, `peanut butter`},
	}
	for k, v := range person {
		fmt.Printf("%v:\n", k)
		for l, v2 := range v {
			fmt.Printf("\t values are: %v and the index position is: %v\n", v2, l)
		}
	}
}
