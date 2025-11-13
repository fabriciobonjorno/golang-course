package main

import "fmt" // Importing fmt package for formatted I/O

func main() {
	salary := map[string]int{
		"ali":   2000,
		"ahmad": 3000,
		"reza":  4000,
	}

	// Iterating over the map and printing each key-value pair
	for name, salary := range salary {
		fmt.Printf("Name: %s, Salary: %d\n", name, salary)
	}

	// Deleting the entry with key "ali"
	delete(salary, "ali")
	fmt.Println("After deleting 'ali':", salary)

	salary["sara"] = 5000 // Adding a new entry to the map
	fmt.Println("After adding 'sara':", salary)

	new_salary := make(map[string]int) // Creating a new empty map
	fmt.Println("New empty map:", new_salary)

	new_salary_2 := map[string]int{} // Creating another new empty map
	fmt.Println("Another new empty map:", new_salary_2)

	// Iterating over the map to print only the values using blank identifier
	for _, salary := range salary {
		fmt.Printf("Salary: %d\n", salary)
	}

}
