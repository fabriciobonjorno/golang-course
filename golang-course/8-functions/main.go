package main

import (
	"errors"
	"fmt"
)

// Importing fmt package for formatted I/O

func main() {
	sum(3, 4)
	fmt.Println("Sum:", sum(3, 4))
	fmt.Printf("---------\n")
	result, ok := sum2(-1, -2)
	fmt.Println("Sum2:", result, "Is Positive?", ok)
	fmt.Printf("---------\n")
	result3, err := sum3(5, -3)
	if err != nil {
		fmt.Println("Error in Sum3:", err)
	} else {
		fmt.Println("Sum3:", result3)
	}

}

func sum(a int, b int) int { // you can use a, b int { instead of a int, b int }
	return a + b
}

// sum2 returns the sum of two integers and a boolean indicating if the sum is positive
func sum2(a, b int) (int, bool) {
	if a+b > 0 {
		return a + b, true
	}
	return a + b, false
}

// sum3 returns the sum of two integers or an error if any integer is negative
func sum3(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("negative numbers are not allowed")
	}
	return a + b, nil
}
