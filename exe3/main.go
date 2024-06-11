package main

import "fmt"

func main() {
	msgvar_44 := sayHello("Alice")
	fmt.Println(msgvar_44)
}

func sayHello(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
