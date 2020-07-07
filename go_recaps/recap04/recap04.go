package main

import (
	"fmt"
)

func main() {
	x := map[string][]string{
		`potato`:  []string{`french fries`, `hash browns`, `latkes`, `pierogis`},
		`beef`:    []string{`burgers`, `meatloaf`, `chili`, `beef broccoli`},
		`chicken`: []string{`chicken parmesean`, `cordon bleu`, `nuggets`, `alfredo`},
		`onion`:   []string{`onion rings`, `omlette`, `pasta`, `pancakes`},
	}
	for k, v := range x {
		fmt.Printf("Popular %v dishes are:\n ", k)
		for l, v2 := range v {
			fmt.Printf("Dishes: %v and their index in the slice is: %v\n", v2, l)
		}
	}
}
