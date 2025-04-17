package main

import (
    "fmt"
)

func main() {
    studentAge := make(map[string]int)
    studentAge["Olivia"] = 16
    studentAge["Isabel"] = 23
    studentAge["Rodrigo"] = 21
    studentAge["Maria"] = 16
    fmt.Println("Olivia's age", studentAge["Olivia"])
    studentGrade := map[string]float64{
        "Clara" : 10, 
        "Olivia" : 9.7,
        "Isabel" : 8.5,
        "Maria" : 9.9,
    }
    for k,v := range studentGrade{
        fmt.Printf("%s got a grade %.1f \n", k, v)
    }
}