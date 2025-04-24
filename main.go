package main

import  "fmt";

func getName() (string, string){
  return "Olivia", "Rodrigo"
}
func main(){
    name, lastname := getName()
    fmt.Println(name);
    fmt.Println(lastname)   

}
