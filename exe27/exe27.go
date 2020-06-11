package main

import "fmt"

func main() {

	type food struct {
		name      string
		quantity  int
		purchased bool
	}
	f1 := food{
		name:      "burger",
		quantity:  2,
		purchased: true,
	}

	f2 := food{
		name:      "fries",
		quantity:  4,
		purchased: false,
	}
	fmt.Println(f1.name, f2.name)
	if f2.purchased == false {
		fmt.Println("sad face!")
	}
}
