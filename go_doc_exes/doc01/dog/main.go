//Package dog converts human years to dog years. Dogs are reportedly 7 times the age of a human
package dog

import "fmt"

//Years will take an integer and multiply it by 7 which converts human to dog years
func Years(n int) (int, error) {
	e := fmt.Errorf("Error: %v is not a valid age", n)
	x := fmt.Println(e)
	var v int

	if n < 0 {
		fmt.Println(e)
	} else {
		v = (n * 7)
	}
	return v, x
}
