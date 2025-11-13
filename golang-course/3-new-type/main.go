package main

import "fmt"

const a = "Hello, World!"

type ID int

// var can be changed at runtime, const cannot
// Global variable declarations
var (
	b bool    = true
	c int     = 42
	d string  = "Hello"
	e float64 = 3.14
	g ID      = 1001 // Using new type ID
)

func main() {
	var f string = "Local variable" // Local variable shadows global 'f' only can be used within main

	// Print all variables
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
