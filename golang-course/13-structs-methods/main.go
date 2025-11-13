package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Adult bool
}

// Method to determine if the person is an adult using a receiver
func (p *Person) isAdult() bool {
	return p.Age >= 18
}

// Another way to implement the isAdult function
// func isAdult(age int) bool {
// 	return age >= 18
// }

func main() {
	p := Person{Name: "Alice", Age: 3}
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Adult:", p.isAdult())
}
