package main

import "fmt"

// dadosPessoa recebe o nome e a idade de uma pessoa e retorna a mesma idade
// e se ela é maior ou menor de idade.
func dadosPessoa(nome string, idade int) (int, string) {
    // Verifica se a idade é maior ou igual a 18
    if idade >= 18 {
        return idade, "Maior de idade"
    }
    return idade, "Menor de idade"
}

func main() {
    // Exemplo de uso da função dadosPessoa
    nome := "Olivia"
    idade := 23

    // Chama a função e recebe a idade e a classificação
    idadeCalculada, classificacao := dadosPessoa(nome, idade)

    // Exibe os resultados
    fmt.Printf("Nome: %s\n", nome)
    fmt.Printf("Idade: %d\n", idadeCalculada)
    fmt.Println(classificacao)
}