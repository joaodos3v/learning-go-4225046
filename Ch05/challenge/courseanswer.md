# Create a more advanced calculator application

## Solution

```go
// calculate() returns the sum of the two parameters
func calculate(input1 string, input2 string, operation string) float64 {
    // Your code goes here.
	var result float64
	value1 := convertInputToValue(input1)
	value2 := convertInputToValue(input2)

	switch operation {
	case "+":
		result = addValues(value1, value2)
	case "-":
		result = subtractValues(value1, value2)
	case "*":
		result = multiplyValues(value1, value2)
	case "/":
		result = divideValues(value1, value2)
	default:
		panic("Invalid operation")
	}
	return result
}

func convertInputToValue(input string) float64 {
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a number", input)
		panic(message)
	}
	return value
}
```

Errors/Improvements:
 - Should use `TrimSpace` method before do `ParseFloat`
 - Could format the panic message
