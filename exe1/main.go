package main

import "fmt"

// var unusedVar2 string

func main() {
	msgvar_22 := sayHello("Alice")
	fmt.Println(msgvar_22)
}

func sayHello(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
