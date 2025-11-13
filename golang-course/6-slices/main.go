package main

import "fmt" // Importing fmt package for formatted I/O

func main() {
	// Creating a slice of strings with initial values
	s := []int{2, 4, 6, 8, 10}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)             // Printing length, capacity, and contents of the slice
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0]) // Printing the slice
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4]) // Printing the slice
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:]) // Printing the slice

	s = append(s, 12) // Appending an element to the slice duplicates the inicial size and capacity

	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2]) // Printing the slice

}
