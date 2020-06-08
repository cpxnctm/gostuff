package main

import (
	"fmt"
)

func main() {
	// create a simple map
	m := map[string]int{ //create the map using a composite literal
		"car":   25,
		"house": 100,
		"pool":  50,
		"food":  30,
	}
	fmt.Println(m)
	fmt.Println(m["food"])
	fmt.Println(m["tv"])
	for k, v := range m {
		fmt.Println(k, v)
		v, ok := m["tv"]
		fmt.Println(v, ok)
		m["something"] = 300 //adding a new entry to the map

		fmt.Println(m)

	}

}
