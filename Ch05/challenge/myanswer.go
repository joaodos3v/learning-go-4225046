package main

import (
	"strconv"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// My code will be here:
type MathOperation func(float64, float64) float64

func createOperationsSlice() map[string]MathOperation {
	var operations = make(map[string]MathOperation)
	operations["+"] = addValues
	operations["-"] = subtractValues
	operations["*"] = multiplyValues
	operations["/"] = divideValues
	return operations

	/**
	Também poderia fazer algo assim e já declarar direto:

	operations := map[string]MathOperation{
		"add":      add,
		"subtract": subtract,
		"multiply": multiply,
	}
	*/
}

func calculate(input1 string, input2 string, operation string) float64 {
	operations := createOperationsSlice()
	value1 := convertInputToValue(input1)
	value2 := convertInputToValue(input2)

	return operations[operation](value1, value2)
}

func convertInputToValue(input string) float64 {
	floatValue, err := strconv.ParseFloat(input, 64)

	if err != nil {
		panic("You need to pass a value that can be converted to float")
	}

	return floatValue
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
