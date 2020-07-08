package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("An error has occurred", err)
	}
	defer f.Close()
	log.SetOutput(f)

	bs, err1 := json.Marshal(p1)
	if err1 != nil {
		fmt.Println("Error:", err1)
	}
	fmt.Println(string(bs))

}
