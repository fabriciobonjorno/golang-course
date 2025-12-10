package main

import "fmt"

// chan<- just receive
// <-chan just send

func receive(name string, hello chan<- string) {
	hello <- name
}

func read(data <-chan string) {
	fmt.Println(<-data)
}

// Thread 1
func main() {

	hello := make(chan string)
	go receive("hello", hello)
	read(hello)
}
