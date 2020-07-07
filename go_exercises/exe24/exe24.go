package main

import "fmt"

func main() {

	//Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
	x := []string{`cats`, `dogs`, `rabbits`, `snakes`}
	y := []string{`cars`, `bikes`, `scooters`, `skates`}
	xy := [][]string{x, y} //multi-dimensional slice
	fmt.Println(xy)
	for k, v := range xy {
		fmt.Printf("Entry: %v\n", k)
		for j, v2 := range v {
			fmt.Printf("At index: %v, the record is: %s\n", j, v2)
		}
	}
}
