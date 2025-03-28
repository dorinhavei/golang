package main

import "fmt" 

func main (){
	a, b := 10, 3
	fmt.Println("The sum is:", a + b)
	fmt.Println("The subtraction is:", a - b)
	fmt.Println("The division is:", a / b)
	fmt.Println("The multiplication is:", a * b)
	fmt.Println("The rest of the division is:", a % b)

	a++
    fmt.Println("Increase a", a)

	if a > 0 && b > 0 {
		fmt.Println("Positive numbers")
	}


}