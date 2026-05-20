package main 

import (
    "fmt"
    "strings"
    "strconv"
    )
    
     var base = map[string]int {
        "(hex)": 16,
        "(bin)": 2,
    }

func main() {
  result := "I have 1E (hex) numbers of files and 10000 (bin) folders on my desktop."
  newResult := strings.Fields(result)
  for i := 0; i < len(newResult); i++ {
  if base, ok := base[newResult[i]]; ok {
       newResults, err := strconv.ParseInt(newResult[i-1], base, 64)
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