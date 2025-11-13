package main

import (
	"fmt"

	"github.com/fabriciobonjorno/go-course/21-module-pack/matematica"
)

func main() {
	sum := matematica.Sum(10, 20)
	fmt.Printf("The result is %v", sum)
}
