package main

import (
	"fmt"

	"github.com/cpxnctm/gostuff/go_doc_exes/doc01/dog"
)

func main() {
	age := 15
	a, b := dog.Years(age)
	if a == 0 {
		fmt.Println(b)
	} else {
		fmt.Printf("Your dog is %v human years old, and that's %v in doggo years.", age, a)
	}

}
