package main

import (
	"fmt"
)

func main() {

	const (
		yr  = 2020
		yr1 = (yr + iota)
		yr2 = (yr + iota)
		yr3 = (yr + iota)
		yr4 = (yr + iota)
	)
	fmt.Println("The current year is:", yr, "and the next four years are:", yr1, ",", yr2, ",", yr3, ",", yr4)

}
