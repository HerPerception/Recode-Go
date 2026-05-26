package main

import (
	"fmt"
	"os"
)

func LastWord() {
	if len(os.Args) != 2 {
		return
	}
	s := os.Args[1]
	if s == " " || s == "" {
		return
	}
	word := []string{}
	i := 0
	n := 0
	for index := len(s) - 1; index >= 0; index-- {
		if index < len(s)-1 && s[index] != ' ' {
			i = index
			break
		}
	}
	for in := 0; in < len(s); in++ {
		if in < i && s[in] == ' ' {
			n = in
		}
	}
	word = append(word, s[n:i+1])
	text := ""
	for _, ch := range word {
		text += string(ch)
	}
	fmt.Println(text)
}
