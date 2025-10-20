package main

import (
	"fmt"
	"time"
)

func main() {
	go say("Hello from the Goroutine!") // turn a function call into a goroutine
	fmt.Println("Hello from main!")

	go func(message string) {
		println(message)
	}("Hello from the anonymous function!")

	time.Sleep(2 * time.Second) // without it, the main thread will finish before than the thread and it'll stop the execution
	fmt.Println("All done!")
}

func say(message string) {
	time.Sleep(1 * time.Second) // "wait a second and then output"
	println(message)
}
