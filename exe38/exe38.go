package main

import "fmt"

// write a callback...

func main() {
	a := []int{1, 2, 3, 4, 5}
	ad := add(a...)
	fmt.Println(ad)
	b := odds(add, ad)
	fmt.Println(b)

}
func add(n ...int) int {
	b := 0
	for _, v := range n {
		b += v
		fmt.Println(v)
	}
	return b
}
func odds(f func(n ...int) int, y ...int) int {
	var xx []int
	for _, v := range y {
		if v%2 != 0 {
			xx = append(xx, v)
		}
	}
	final := f(xx...)
	return final
}
