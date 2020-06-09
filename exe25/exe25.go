package main

import (
	"fmt"
)

func main() {
	//create a mapping between an object and corresponding information such as their favorite things. print out the values as well as their position within the slice

	x := map[string][]string{
		`cats`: []string{`meow`, `treats`, `lasers`, `sleep`},
		'dogs': []string{`woof`, 'bones','fetching','pets'},
		'fish': []string{`swimming`, `breathing`,`watching`,`floating`}
		fmt.Println(x)

	}
}
