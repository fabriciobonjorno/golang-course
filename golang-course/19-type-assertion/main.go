package main

import "fmt"

func main() {
	var myVar interface{} = "Hello, World!"
	println(myVar.(string))

	result, ok := myVar.(int)
	fmt.Printf("Result value is %v and ok result is %v", result, ok)
}
