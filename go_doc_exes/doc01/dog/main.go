//Package dog converts human years to dog years. Dogs are reportedly 7 times the age of a human
package dog

//Years will take an integer and multiply it by 7 which converts human to dog years.
//If the provided integer is negative, then Years will return an error string
func Years(n int) (int, string) {
	var a int
	var b string
	if n > 0 {
		a = (n * 7)
	} else {
		b = "Invalid age provided. Ages cannot be negative."
	}
	return a, b

}
