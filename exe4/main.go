package main

import "fmt"

var unusedVar string

func main() {
	msg_var := sayHello("Alice")
	fmt.Println(msg_var)
}

func sayHello(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
