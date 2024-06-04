package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	// Unused variable
	var unusedVar string

	// Ineffectual assignment
	x := 42
	x = 43
	y = 90

	// Redundant import
	_ = fmt.Sprintf("This import is redundant")

	// Unnecessary conversion
	y := int64(123)
	fmt.Println(int64(y))
}
