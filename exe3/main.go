package main

import "fmt"

var unusedVar string

func main() {
	msg_var44 := sayHello("Alice")
	fmt.Println(msg_var44)
}

func sayHello(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
