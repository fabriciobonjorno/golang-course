package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Adult bool
}

func isAdult(age int) bool {
	return age >= 18
}

func main() {
	p := Person{Name: "Alice", Age: 3}
	p.Adult = isAdult(p.Age)
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Adult:", p.Adult)
}
