package main

import "fmt"

func main(){
  scores := make([]int, 0, 5)
  c := cap(scores)
  fmt.Println("Capacité", c)

  for i := 0; i < 25; i++ {
    scores = append(scores, i)
  }

  c = cap(scores)
  fmt.Println("Nouvelle capacité", c)
}
