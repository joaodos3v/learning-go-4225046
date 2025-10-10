package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(("Enter text: "))
	str, _ := reader.ReadString('\n') // _ is used to ignore the error
	fmt.Println(str)

	fmt.Print(("Enter a number: "))
	str, _ = reader.ReadString('\n')                         // We should only use := when we're declaring the variables for the first time
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64) // 64 is the precision

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The number:", f)
	}
}
