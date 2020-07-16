package main

import (
	"fmt"

	"github.com/cpxnctm/gostuff/go_testing/testing04/quote"
	"github.com/cpxnctm/gostuff/go_testing/testing04/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
