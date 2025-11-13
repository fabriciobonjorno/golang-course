package main

import "fmt" // Importing fmt package for formatted I/O

// const cannot be changed at runtime
// Global constant declaration

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
	fmt.Printf("B type is %T\n", b)
	fmt.Printf("C type is %T\n", c)
	fmt.Printf("D type is %T\n", d)
	fmt.Printf("E type is %T\n", e)
	fmt.Printf("F type is %T\n", f)
	fmt.Printf("G type is %T\n", g)
	fmt.Printf("A type is %T\n", a)

}
