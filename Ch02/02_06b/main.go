package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t)

	now := time.Now()
	fmt.Printf("The time currently is %s\n", now)
	fmt.Printf("This object's type is %T\n", now)

	fmt.Printf(now.Format(time.ANSIC) + "\n")

	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf(tomorrow.Format(time.ANSIC) + "\n")

	format := "Mon 2006-02-01"
	fmt.Printf(tomorrow.Format(format) + "\n")
}
