package main

import "fmt"

func main() {
    msg := sayHello("Alice")
	msg2 := sayHello2("say hello 2")
    fmt.Println(msg)
	fmt.Println(msg2)
}

func sayHello(name string) string {
    return fmt.Sprintf("Hi %s", name)
}

func sayHello2(name string) string {
    return fmt.Sprintf("Hi %s", name)
}