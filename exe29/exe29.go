package main

import (
	"fmt"
)

func main() {
	type vehicle struct {
		doors int
		color string
	}
	type truck struct {
		vehicle
		fourWheel bool
	}
	type sedan struct {
		vehicle
		luxury bool
	}
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: `green`,
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: `tan`,
		},
		luxury: true,
	}
	fmt.Println(t.doors)
	fmt.Println(s.luxury)

}
