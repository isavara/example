package main

import "fmt"

var unusedVar1 string

func main() {
	msgvar22 := sayHello("Alice")
	fmt.Println(msgvar22)
}

func sayHello(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
