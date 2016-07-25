package main

import "fmt"

func main(){
  // Tableaux #1
  var scores [10]int
  scores[0] = 50

  // Tableaux #2
  score := [4]int{50, 105, 512, 98}


  fmt.Println(scores[0], score[2])
  for index, value := range score {
    fmt.Println(index, value)
  }
}
