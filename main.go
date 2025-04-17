package main

import (
    "fmt"
)
func division(dividend int, divider int) (int, string){
    if divider == 0 {
    return 0, "error in dividion by 0"
    }
    return dividend / divider, ""
}
func basicOperation(a int, b int) (int, int, int){
    sum := a + b
    multiplication := a * b
    subtraction := a - b
    return sum, multiplication, subtraction
}

func main() {
 result, erro := division(10,2)
 if erro != "No error" {
    fmt.Println(erro)
 } else {
    fmt.Println("Result of division is:", result, erro)
 }
  sum, mult, sub := basicOperation(10,2)
  fmt.Println(sum)
  fmt.Println(mult)
  fmt.Println(sub)
}