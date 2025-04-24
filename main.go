package main

import  "fmt";

func analiseGrades(note1, note2 float64) (float64, string){
    media := (note1 + note2) / 2
    if media < 6 {
    return media, "Failed"
    } else { 
        return media, "Approved"
    }
    
}

func main(){
    media, result := analiseGrades(7.5, 5.5)
     fmt.Println("Media:", media)
     fmt.Println("Result:", result)
    

}


