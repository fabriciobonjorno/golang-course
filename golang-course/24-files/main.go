package main

import (
	"bufio"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	// You can use either WriteString or Write method to write to the file
	// size, err := f.WriteString("Hello, World!")
	size, err := f.Write([]byte("Hello, World!"))
	if err != nil {
		panic(err)
	}

	println("Size is", size)

	defer f.Close()

	// Read the file content
	content, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	println("File content:", string(content))

	// read line by line
	lines, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	// using bufio to read line by line
	reader := bufio.NewReader(lines)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		println("Read bytes:", n, "Content:", string(buffer[:n]))
	}
	defer lines.Close()
}
