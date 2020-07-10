package main

import (
	"fmt"

	"github.com/cpxnctm/gostuff/go_doc_exes/doc01/dog"
)

func main() {
	a, b := dog.Years(-7)
	if a == 0 {
		fmt.Println(b)
	} else {
		fmt.Println(a)
	}

}
