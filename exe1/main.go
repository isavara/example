package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	// Unused variable

	// Ineffectual assignment
	var x = 42

	// Redundant import
	_ = fmt.Sprintf("This import is redundant")

	// Unnecessary conversion
	y := int64(123)
	fmt.Println(int64(y))
}
