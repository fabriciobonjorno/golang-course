package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Adult bool
	Adress
}

type Adress struct {
	City   string
	Street string
	Number int
}

func isAdult(age int) bool {
	return age >= 18
}

func main() {
	p := Person{Name: "Alice", Age: 3}
	p.Adult = isAdult(p.Age)
	p.Adress = Adress{City: "Wonderland", Street: "Main St", Number: 123}
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Adult:", p.Adult)
	fmt.Println("City:", p.City)
	fmt.Println("Street:", p.Street)
	fmt.Println("Number:", p.Number)

}
