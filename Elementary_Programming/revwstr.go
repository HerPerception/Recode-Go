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
	end := len(word)
	text := ""
	newText := ""
	for i := 0; i < len(word); i++ {
		if word[i] == ' ' {
			continue
		} else if word[i] != ' ' {
			start = i
			break
		}
	}
	for i := len(word) - 1; i >= 0; i-- {
		if word[i] != ' ' {
			break
		}
		end = i
	}
	text = (word[start:end])
	slice := []string{}
	current := ""
	for i := 0; i < len(text); i++ {
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
	for i := len(slice) - 1; i >= 0; i-- {
		newSlice = append(newSlice, slice[i])
	}
	for ch := range newSlice {
		if ch == len(newSlice)-1 {
			newText += newSlice[ch]
		} else {
			newText += newSlice[ch] + " "
		}
	}
	fmt.Printf("%q\n", newText)
}
