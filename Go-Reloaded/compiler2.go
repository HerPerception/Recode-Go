package main 

import (
    "fmt"
    "strings"
    "strconv"
    )

func main() {
  result := "I have 1E (hex) numbers of files and 10 (bin) folders on my desktop."
  newResult := strings.Split(result, " ")
  for i := 0; i < len(newResult); i++ {
    switch newResult[i] {
        case "(hex)":
       newResults, err := strconv.ParseInt(newResult[i-1], 16, 64)
       if err != nil {
           fmt.Println("Error parsing number.")
           continue
       }
             newResult[i-1] = fmt.Sprintf("%d", newResults)
                   newResult = append(newResult[:i], newResult[i+1:]...)
                   i--
       case "(bin)":
       newResults, err := strconv.ParseInt(newResult[i-1], 2, 64)
       if err != nil {
           fmt.Println("Error parsing number.")
           continue
       }
             newResult[i-1] = fmt.Sprintf("%d", newResults)
                   newResult = append(newResult[:i], newResult[i+1:]...)
                   i--
  }
  }
  final := strings.Join(newResult, " ")
  fmt.Println(final)
}