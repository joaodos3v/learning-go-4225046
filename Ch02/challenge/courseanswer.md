```go
func calculate(value1 string, value2 string) float64 {
	float1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}
	float2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}
	return float1 + float2
}
```

Errors:
  - Forgot `strings.TrimSpace(str)`
	- Forgot to use `panic` to stop the code