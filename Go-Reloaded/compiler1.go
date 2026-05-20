package main 

import (
    "fmt"
    "strings"
    "strconv"
    )

func main() {
  result := "I have 42 (hex) numbers of files and 1e (hex) folders on my desktop."
  newResult := strings.Split(result, " ")
  for i := 0; i < len(newResult); i++ {
    if newResult[i] == "(hex)"{
       newResults, err := strconv.ParseInt(newResult[i-1], 16, 64)
       if err != nil {
           fmt.Println("Error parsing number.")
           break
       }
           newResult[i-1] = fmt.Sprintf("%d", newResults)
                   newResult = append(newResult[:i], newResult[i+1:]...)
                   i--
  }
  }
  final := strings.Join(newResult, " ")
  fmt.Println(final)
}