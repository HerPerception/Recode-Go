package main

import (
	"fmt"
	"os"
)

func CleanStr() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		fmt.Println()
		return
	}
	arg := os.Args[1]
	text := []string{}
	start := 0
	end := 0
	word := ""
	newText := ""
	for index := 0; index < len(arg); index++ {
		if arg[index] != ' ' {
			break
		}
		start = index
	}
	for index := len(arg) - 1; index >= 0; index-- {
		if arg[index] != ' ' {
			break
		}
		end = index
	}
	text = append(text, arg[start+1:end])

	for _, ch := range text {
		word += ch
	}
	for i := range word {
		if word[i] == ' ' && word[i+1] == ' ' {
			continue
		} else {
			newText += string(word[i])
		}
	}
	fmt.Printf("%q\n", newText)
}
