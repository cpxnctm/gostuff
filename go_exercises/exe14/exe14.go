package main

import "fmt"

func main() {
	i := 1988 //init
	for {
		fmt.Println(i)
		if i > 2020 { //conditional
			break
		}
		i++ //post
	}
}
