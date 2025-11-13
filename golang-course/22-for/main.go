package main

func main() {

	// Print numbers from 1 to 10 using a for loop
	for i := 1; i <= 10; i++ {
		println(i)
	}

	numbers := []int{1, 2, 3, 4, 5}

	// Iterate over the slice using range and print each number
	for k, v := range numbers {
		println(k, v)
	}

	u := 0
	// Use a for loop to mimic a while loop
	for u < 5 {
		println(u)
		u++
	}
}
