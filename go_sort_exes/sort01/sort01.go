package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{3, 7, 20, 100, 55555, 8, 66, 85, 12, 49, 78}
	sort.Ints(a)
	fmt.Println(a)
}
