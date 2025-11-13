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
	// When array is declared you can´t change its size
	var array [5]int = [5]int{1, 2, 3, 4, 5} // Array declaration and initialization 5 positions with int type
	fmt.Println("Array:", array)
	fmt.Println(len(array))

	// Accessing array elements using index whith traditional for loop

	for i := 0; i < len(array); i++ {
		fmt.Printf("Element at index %d is %d\n", i, array[i])
	}

	fmt.Println("------------")
	// Accessing array elements using range-based for loop

	for i, value := range array {
		fmt.Printf("Element at index %d is %d\n", i, value)
	}

}
