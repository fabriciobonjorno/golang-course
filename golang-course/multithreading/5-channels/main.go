package main

import "fmt"

// Thread 1
func main() {
	channel := make(chan string)

	// Thread 2
	go func() {
		channel <- "Hello World"
	}()

	msg := <-channel
	fmt.Println(msg)
}
