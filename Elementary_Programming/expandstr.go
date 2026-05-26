package main

import (
	"fmt"
	"os"
)

func ExpandStr() {
	if len(os.Args) != 2 || len(os.Args[1]) == 0 {
		fmt.Println()
		return
	}
	text := os.Args[1]
	word := ""
	start := 0
	end := len(text)
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			continue
		} else if text[i] != ' ' {
			start = i
			break
		}
	}
	for i := len(text) - 1; i >= 0; i-- {
		if text[i] != ' ' {
			break
		}
		end = i
	}
	//slice := []string{}
	word = (text[start:end])

	// for _, ch := range slice {
	// 	word += ch
	// }
	newText := ""
	for i := range word {
		if word[i] == ' ' && word[i-1] != ' ' {
			for n := 0; n < 3; n++ {
				newText += string(word[i])
			}
		} else if word[i] == ' ' && word[i-1] == ' ' {
			continue
		} else {
			newText += string(word[i])
		}
	}
	fmt.Printf("%q\n", newText)
}
