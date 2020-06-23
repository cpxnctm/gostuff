package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pwd := `somethingsomething!`
	fmt.Println(pwd)
	spwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(spwd)

	err = bcrypt.CompareHashAndPassword(spwd, []byte(pwd))
	if err != nil {
		fmt.Println("The passwords do not match. Please try again")
	} else {
		fmt.Println("Authentication Successful")
	}

}
