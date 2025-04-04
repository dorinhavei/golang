package main

import (
  "strings"
  "fmt"
  "sort"
)
func main() {
   greeting := "Hello guys!"
   fmt.Println(strings.Contains(greeting, "friends"))
   fmt.Println(strings.ReplaceAll(greeting,"Hello", "Hi"))
   fmt.Println(strings.ToUpper(greeting))
   fmt.Println(strings.Index(greeting, "my"))
   fmt.Println(strings.Split(greeting, " "))
   ages := []int{50, 80, 10}
   sort.Ints(ages)
   fmt.Println(ages)
   index := sort.SearchInts(ages, 50)
   fmt.Println(index)
   names := []string{"Olivia", "Isabel", "Rodrigo"}
   sort.Strings(names)
   fmt.Println(names)
   fmt.Println(sort.SearchStrings(names, "Olivia"))

   x:=46

   for x < 5 {
	fmt.Println(x)
	x++
   }

   for i:=0; i <5; i++{
	fmt.Println("for 2:", i)
   }

   for i:=0; i <len(names); i++{
	fmt.Println(names[i])
   }

   for index, value := range names{
	fmt.Println("O índice é", index, "e o valor é", value)
   }

   for index, value:= range ages{
	fmt.Println("O índice é", index, "e o valor é", value)
   }

   superhero:=[]string{"Spiderman", "Ironman", "Batman"}

   for index, value := range superhero{
	fmt.Println("The hero's number is:", index, "and his name is:", value)
   }
}