package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World!"
	showTypeAndValue(x)
	showTypeAndValue(y)
}

func showTypeAndValue(t interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", t, t)
}
