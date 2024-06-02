package main

import "fmt"

func main() {
    msg_var := sayHello("Alice")
    fmt.Println(msg_var)
}

func sayHello(name string) string {
    return fmt.Sprintf("Hi %s", name)
}
