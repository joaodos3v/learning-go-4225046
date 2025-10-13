package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)

	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"

	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "New York"
	fmt.Println(states)

	for key, value := range states {
		fmt.Printf("%v: %v\n", key, value)
	}

	keys := make([]string, len(states))
	i := 0
	for key := range states {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	fmt.Println("\nSorted")

	for i := range keys {
		fmt.Println((states[keys[i]]))
	}
}
