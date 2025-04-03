package main

import "fmt" 

func main (){
	var ages = [4]int{17,16,20,40}
	names:= [4]string{"Olivia", "Isabel", "Rodrigo", "Liv"}
    fmt.Println(ages)
	fmt.Println(names)
	names[3] = "Rogério"
	fmt.Println(names)
}