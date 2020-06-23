package main

import (
	"encoding/json"
	"fmt"
)

type coffees struct {
	Name     string `json:"Name"`
	Location string `json:"Location"`
	Flavor   bool   `json:"Flavor"`
	Cost     int    `json:"Cost"`
}

func main() {
	c := `[{"Name":"regular","Location":"america","Flavor":false,"Cost":10},{"Name":"decaf","Location":"s. america","Flavor":false,"Cost":12},{"Name":"fancy","Location":"africa","Flavor":true,"Cost":15}]`
	xc := []byte(c)
	drinks := []coffees{}
	fmt.Printf("The coffees are:\n")
	err := json.Unmarshal(xc, &drinks)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	for _, v := range drinks {
		fmt.Println(v.Name, v.Location, v.Flavor, v.Cost)
	}
}
