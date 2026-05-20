package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	if len(os.Args[1]) == 0 || os.Args[1] == " " {
		fmt.Println()
		return
	}
	word := os.Args[1]
	start := 0
	for i := 0; i < len(word); i++ {
		if word[i] == ' ' {
			continue
		} else {
			start = i
			break
		}
	}
	end := len(word)
	for i := len(word) - 1; i >= 0; i-- {
		if word[i] != ' ' {
			break
		}
		end = i
	}

	text := word[start:end]
	//snewText := ""
	// for i := range text {
	// 	if i == 0 {
	// 		newText += string(text[len(text)-1])
	// 	} else if i == len(text)-1 {
	// 		newText += string(text[0])
	// 	} else if i > 0 && text[i] == ' ' && text[i-1] == ' ' {
	// 		continue
	// 	} else {
	// 		newText += string(text[i])
	// 	}
	// }
	current := ""
	slice := []string{}
	for i := range text {
		if text[i] != ' ' {
			current += string(text[i])
		}
		if text[i] == ' ' && current != "" {
			slice = append(slice, current)
			current = ""
		}
	}
	if current != "" {
		slice = append(slice, current)
	}
	newSlice := []string{}
	for i := 1; i < len(slice); i++ {
		newSlice = append(newSlice, slice[i])
	}
	newSlice = append(newSlice, slice[0])
	newText := ""
	for i := range newSlice {
		if i == len(newSlice)-1 {
			newText += string(newSlice[i])
		} else {
			newText += string(newSlice[i]) + " "
		}
	}
	fmt.Printf("%q\n", newText)
}
