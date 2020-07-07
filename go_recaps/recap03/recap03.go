package main

import "fmt"

func main() {

	x := map[string][]string{
		`summer`: []string{`hot`, `sunshine`, `sports`, `ice cream`},
		`fall`:   []string{`cool`, `leaves`, `long days`, `halloween`},
		`winter`: []string{`cold`, `skiing`, `snow`, `ice`},
		`spring`: []string{`nice`, `rain`, `outdoors`, `patio weather`},
	}
	for k, v := range x { // here, k is the key and v is the value
		fmt.Printf("The season is: %v\n", k)
		for l, v2 := range v { // because the range is equal to v (which is the value), l is the index and v2 is the mapped value
			fmt.Printf("It is: %v and its index is: %v\n", v2, l)
		}
	}

}
