package main

import "fmt"

var unusedVar string

func main() {
	msg_var22 := sayHello("Alice")
	fmt.Println(msg_var22)
}

func sayHello(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
