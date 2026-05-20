package main

import "fmt"

// func lastLetter(s string) string {
// 	if len(s) == 0 {
// 		return ""
// 	}
// 	return string(s[len(s)-1])
// }

func main(){
	word := ""
	str := ""
	var operator string
	for {
	fmt.Print("Enter a word: ")
	fmt.Scanln(&word)
	fmt.Println("Enter desired operation(lastLetter, capitalize, delete): ")
	fmt.Scanln(&operator)
	switch operator {
	case "capitalize": 
		for _, ch := range word {
			if ch >= 'a' && ch <= 'z' {
				ch -= 32
				str += string(ch)
			} else {
				str += string(ch)
			}
		}
		case "lastLetter":
			fmt.Print(string(word[len(word)-1]))
	
case "delete":
	index := 0
	fmt.Print("Enter index to delete: ")
	fmt.Scanln(&index)
	in := len(word)-1

		if index > len(word)-1 {
		fmt.Printf("Index out of range, the range of indexes is from 0 to %d\n",in )
   }

	for idx, ch := range word {
		if index <= len(word)-1 && idx == index {
			continue
		} else {
			str += string(ch)
		}
	}
  }
 }
	fmt.Println(str)

}

