package main

import (
    "fmt"
)

func main() {
    var numeros [5]int
    var soma int

    // Solicita ao usuário que digite os 5 valores
    fmt.Println("Digite 5 números inteiros:")

    for i := 0; i < 5; i++ {
        fmt.Printf("Número %d: ", i+1)
        fmt.Scanln(&numeros[i])
        soma += numeros[i]
    }

    // Exibe a soma dos elementos
    fmt.Printf("A soma dos números digitados é: %d\n", soma)
}