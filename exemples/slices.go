package main

import "fmt"

func main(){
  // Déclaration #1
  //scores := []int{1,4,293,4,9}

  // Déclaration #2
  //scores := make([]int, 10)
  scores := make([]int, 0, 1)
  scores = append(scores, 5)
  scores = append(scores, 5)
  scores = append(scores, 5)
  scores = append(scores, 5)
  scores = append(scores, 5)
  //scores = scores[0:10]
  //scores[7] = 9000
  fmt.Println(scores)

  for index, value := range scores {
    fmt.Println(index, value)
  }
}
