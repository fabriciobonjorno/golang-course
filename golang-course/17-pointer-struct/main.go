package main

import "fmt"

type Customer struct {
	Name  string
	Email string
}

// Method with pointer receiver to modify the struct's Name field
func (c *Customer) walk() {
	c.Name = "Bonjorno"
	fmt.Printf("Customer %v is walking\n", c.Name)
}

func main() {
	customer := Customer{Name: "John Doe", Email: "john@example.com"}
	customer.walk()
	fmt.Printf("Struct value  is %v", customer.Name)
}
