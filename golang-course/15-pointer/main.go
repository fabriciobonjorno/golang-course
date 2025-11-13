package main

import "fmt"

func main() {
	// Declare an integer variable pointed to by a pointer
	a := 42
	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)

	// Declare a pointer variable that holds the address of a
	var p *int = &a
	fmt.Println("Value of p (address of a):", p)
	fmt.Println("Value pointed to by p:", *p)

	*p = 100
	fmt.Println("New value of a after modifying through p:", a)
}
