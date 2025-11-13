package main

// Before generics, we had to write separate functions for each type.
// Here are two functions that sum values in a map for int and float64 types.
// func IntegerSum(m map[string]int) int {
// 	total := 0
// 	for _, v := range m {
// 		total += v
// 	}
// 	return total
// }

// func FloatSum(m map[string]float64) float64 {
// 	total := 0.0
// 	for _, v := range m {
// 		total += v
// 	}
// 	return total
// }

// We can define a type constraint using a named interface.
type Number interface {
	~int | ~float64
}

// With generics, we can create a single function that works for multiple types.

func Sum[T Number](m map[string]T) T {
	var total T
	for _, v := range m {
		total += v
	}
	return total
}

// Alternatively, we can define the type constraint inline without a named interface.
// func Sum[T int | float64](m map[string]T) T {
// 	var total T
// 	for _, v := range m {
// 		total += v
// 	}
// 	return total
// }

func main() {

	// Using the generic Sum function for int map
	myMap := map[string]int{"a": 1, "b": 2, "c": 3}
	result := Sum(myMap)
	println(result)

	// Using the generic Sum function for float64 map
	myMapFloat := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
	resultFloat := Sum(myMapFloat)
	println(resultFloat)

	// The above commented code can be uncommented to see the results.

	// For demonstration purposes, here's how you would call the functions:

	// myMap := map[string]int{"a": 1, "b": 2, "c": 3}
	// result := IntegerSum(myMap)
	// println(result)

	// myMapFloat := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
	// resultFloat := FloatSum(myMapFloat)
	// println(resultFloat)
}
