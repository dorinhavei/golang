package main

import (
    "fmt"
)

func main() {
    var balance float32
    fmt.Println("enter the initial value:")
    fmt.Scan(&balance)
    escolherOperacao(&balance)
}

func escolherOperacao(balance *float32) {
    var option int
    fmt.Println("Do you want to withdraw (1), deposit (2) or close the session (3)?")
    fmt.Scan(&option)
    switch option{ 
 
    case 1:
        sacar(balance)
    case 2:
        depositar(balance)
    case 3:
        encerrar()
    default:
        fmt.Println("Invalid option")
    }
}

func depositar(balance *float32) {
    var value float32
    fmt.Println("what amount do you want to deposit?")
    fmt.Scan(&value)
    *balance += value
    fmt.Println("your current balance:", *balance)
}

func sacar(balance *float32) {
    var value float32
    fmt.Println("How much do you want to withdraw?")
    fmt.Scan(&value)
    if value > *balance {
        fmt.Println("insufficient balance")
    } else {
        *balance -= value
        fmt.Println("Syour current balance:", *balance)
    }
}

func encerrar() {
    fmt.Println("Session closed")
}