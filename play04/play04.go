package main

import (
	"fmt"
)

func main() {
	a, b := foo("What day is it?", "Friday!")
	fmt.Println(a, b)
}
func foo(s string, t string) (string, string) {
	return "", fmt.Sprint("Is it the end of the week again? ", s, " ", t)
}
