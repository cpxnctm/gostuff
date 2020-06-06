package main

import (
	"fmt"
)

//Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.

func main() {
	favSport := "football"
	switch favSport {
	case ("cooking"):
		fmt.Println("i like to cook")
	case ("driving"):
		fmt.Println("driving is fun")
	case ("football"):
		fmt.Println("football is fun to watch")
		fallthrough
	case ("baseball"):
		fmt.Println("this prints because of fallthrough")
	}
}
