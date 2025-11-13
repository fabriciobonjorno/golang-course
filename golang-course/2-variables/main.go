package main

import "fmt"

const a = "Hello, World!"

// var can be changed at runtime, const cannot
// Global variable declarations
var (
	b bool    = true
	c int     = 42
	d string  = "Hello"
	e float64 = 3.14
)

func main() {
	var f string = "Local variable" // Local variable shadows global 'f' only can be used within main
	// short hand variable declaration var f := "Local variable"

	// Print all variables
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
