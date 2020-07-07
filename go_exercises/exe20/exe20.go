package main

import (
	"fmt"
)

func main() {

	//Using a COMPOSITE LITERAL: create a SLICE of TYPE int and assign 10 values

	a := []int{42, 43, 44, 45, 46}
	b := []int{47, 48, 49, 50, 51}
	c := []int{44, 45, 46, 47, 48}
	d := []int{43, 44, 45, 46, 47}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	//for r := range a {
	//fmt.Printf("%v", a[r]) // range over the slice and print out the values
	//fmt.Printf("%T\n", a[r]) //range over the slice and print out the type
	//for i := range b {
	//fmt.Printf("%v", b[i]) // range over the slice and print out the values
	//fmt.Printf("%T\n", b[i])
	//for v := range c {
	//fmt.Printf("%v", c[v]) // range over the slice and print out the values
	//fmt.Printf("%T\n", c[v])
	//for n := range d {
	//fmt.Printf("%v", d[n]) // range over the slice and print out the values
	//	fmt.Printf("%T\n", d[n])

	//}
	//}
	//}
	//}
}
