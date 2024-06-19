package main

import "fmt"

// var unusedVar2 string

func main() {
	msgvar2 := sayHello("Alice")
	fmt.Println(msgvar2)
}

func sayHello(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
