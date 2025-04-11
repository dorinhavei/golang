package main

import (
    "fmt"
)

func sayGreeting(name string) {
    fmt.Println("Hello!", name)
    }
func addNumber(number1 int, number2 int) int{
    return number1 + number2
}
func main() {
    sayGreeting("Olivia")
    result := addNumber(10,20)
    fmt.Println(result)
}