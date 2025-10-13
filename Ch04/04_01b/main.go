package main

import (
	"fmt"
)

func main() {
	theAnswer := 42 // See the comment below
	var result string

	if theAnswer < 0 { // We can define a scoped-variable like this ---> if theAnswer := 42; theAnswer < 0 {
		result = "Less than zero"
	} else if theAnswer == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println(result)
}
