package word

import (
	"strings"
)

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count takes in a string and returns the number of words separated by whitespace
func Count(s string) int {
	// write the code for this func
	st := strings.Fields(s)
	ct := 0
	//var m2 []int

	for i := range st {
		ct = i
		//m2[i]++
	}
	return ct
}
