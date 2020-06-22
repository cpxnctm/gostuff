package main

import (
	"encoding/json"
	"fmt"
)

type pet struct {
	Name   string `json:"Name"`
	Animal string `json:"Animal"`
	Age    int    `json:"Age"`
	Cute   bool   `json:"Cute"`
}

func main() {
	p := `[{"Name":"Bootsy","Animal":"Kitty","Age":9,"Cute":true},{"Name":"Beta","Animal":"Pupper","Age":1,"Cute":true}]`
	xp := []byte(p)

	animals := []pet{}
	//fmt.Printf("%T\n", animals)

	err := json.Unmarshal(xp, &animals)
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, v := range animals {
		//fmt.Println("Location: ", k)
		fmt.Println(v.Name, v.Animal, v.Age, v.Cute)
	}

}
