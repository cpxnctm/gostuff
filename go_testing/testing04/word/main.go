package word

import (
	"strings"
)

//UseCount takes in a string and returns the number of occurrences for each word
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

// Count takes in a string and returns the number of total words separated by whitespace
func Count(s string) int {
	st := strings.Fields(s)
	ct := 0
	for i := range st {
		ct = i
	}
	return ct
}

// Better should be a more efficient word counting function than Count
func Better(s string) int{
	st := strings.Fields(s)
	return len(st) -1
}

