package main

func main() {
	switch day := "Saturday"; day {
	case "Monday":
		println("It's Monday!")
	case "Saturday", "Sunday":
		println("It's the weekend!")
	default:
		println("It's a weekday.")
	}

	if num := 10; num%2 == 0 {
		println("The number is even.")
	} else {
		println("The number is odd.")
	}
}
