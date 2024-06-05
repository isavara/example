package main

import "fmt"

func main() {
	msgvar44 := sayHello("Alice")
	fmt.Println(msgvar44)
}

func sayHello(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
