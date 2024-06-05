package main

import "fmt"

var unusedVar string

func main() {
	msgvar := sayHello("Alice")
	fmt.Println(msgvar)
}

func sayHello(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
