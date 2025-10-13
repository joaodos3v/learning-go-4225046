```go
func slicesToObjects(colorNames []string, hexValues []int) []Color {
	colors := make([]Color, 0, len(colorNames))
	for i := 0; i < len(colorNames); i++ {
		colors = append(colors, Color{colorNames[i], hexValues[i]})
	}
	return colors
}
```

Errors:
 - Forgot (misconfused) to use the third parameter for `make` function