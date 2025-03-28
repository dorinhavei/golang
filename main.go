package main

import "fmt" 

func main (){
	var number1 int
	var number2 int
	fmt.Print("Type one number:")
	fmt.Scan(&number1)
	fmt.Print("type another number:")
	fmt.Scan(&number2)

	fmt.Println("The sum is:", number1 + number2)
	fmt.Println("The subtraction is:", number1 - number2)
	fmt.Println("The division is:", number1 / number2)
	fmt.Println("The multiplication is:", number1 * number2)
	fmt.Println("The rest of the division is:", number1 % number2)

	}