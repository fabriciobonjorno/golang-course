package main

import "fmt"

// Importing fmt package for formatted I/O

func main() {
	// Using an anonymous function to calculate the total sum multiplied by 2
	// closure captures the sum function
	total := func() int {
		return sum(1, 2, 3, 4, 5) * 2
	}()

	// Example usage of the sum function
	result := total
	fmt.Println("The sum is:", result)
}

// sum takes a variable number of integer arguments and returns their sum.
func sum(numbers ...int) int {
	total := 0
	// Iterate over the numbers and accumulate the total
	for _, n := range numbers {
		total += n
	}
	return total
}
