package main

import "fmt"

type customErr struct {
	warning string
}

func (w customErr) Error() string {
	return fmt.Sprintf("The error is: %v\n", w.warning)
}

func main() {
	x := customErr{
		warning: "The temperatures are going to be high today. Be careful.",
	}
	foo(x)
}

func foo(e error) {
	fmt.Println("Foo is running!", e)
}
