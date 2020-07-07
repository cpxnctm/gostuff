package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	_, err := os.Open("somefile.txt")
	if err != nil {
		log.Println("We had an error:", err)
	}
	_, err = os.Open("somefile.txt")
	if err != nil {
		fmt.Println("We had an error:", err)
	}
}
