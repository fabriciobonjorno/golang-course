package main

import "fmt"

// sum takes two integer pointers, modifies the value at the first pointer,
// and returns the sum of the values pointed to by both pointers.
func sum(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	x := 10
	y := 20
	result := sum(&x, &y)
	fmt.Println("Sum:", result)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}
