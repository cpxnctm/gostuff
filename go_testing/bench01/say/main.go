//Package say prints out a string that is passed into it and adds the string:
// Have a great day
package say

import "fmt"

//Kind takes in a string and adds: Have a great day
func Kind(s string) string {
	return fmt.Sprint("Have a great day,", s)
}
