// Write your answer here, and then test your code.
package main

import (
	"fmt"
	"strconv"
)

// Uncomment this import section to use required Go packages

// "strings"

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// calculate() returns the the result of adding 2 numbers
// after parsing them from strings
func calculate(value1 string, value2 string) float64 {
	// Convert the strings to a float64
	fNumber1, err1 := strconv.ParseFloat(value1, 64)
	fNumber2, err2 := strconv.ParseFloat(value2, 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Both inputs should be numbers!")
	}

	// Calculate and return the result
	sum := fNumber1 + fNumber2

	return sum
}
