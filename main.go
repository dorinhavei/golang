package main

import "fmt" 

func main (){
	var user string
	var password string

	fmt.Print("Enter your username:")
	fmt.Scan(&user)
	fmt.Print("Type your password:")
	fmt.Scan(&password) 

	if user != "admin" && password != "1234" {
	fmt.Println("Access denied") 
    }else{
	fmt.Println("Access allowed") 
   }
}