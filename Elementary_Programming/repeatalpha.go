package main

import (
	"fmt"
	"os"
)

func RepeatAlpha() {
	if len(os.Args) != 2 {
		return
	}
	s := os.Args[1]
	isLetter := false
	word := ""
	index := 0
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
			isLetter = true
		}
		if isLetter == true && (ch >= 'A' && ch <= 'Z') {
			index = int(ch) - 64
			for index > 0 {
				word += string(ch)
				index -= 1
			}
		} else if isLetter == true && (ch >= 'a' && ch <= 'z') {
			index = int(ch) - 96
			for index > 0 {
				word += string(ch)
				index -= 1
			}
		} else {
			word += string(ch)
		}

	}
	fmt.Println(word)
}
