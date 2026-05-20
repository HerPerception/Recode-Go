package main

import (
	"fmt"
	"strings"
)

func FirstWord(s string) string {
	s = strings.TrimSpace(s)
	word := ""
	for _, ch := range s {
		word += string(ch)
		if ch == ' ' {
			break
		}
	}
	return word
}

func LastWord(s string) string {
	s = strings.TrimSpace(s)
	word := ""
	index := 0
	for i := range s {
		if s[i] == ' ' {
			index = i + 1
		}
	}
	for i := index; i < len(s); i++ {
		word += string(s[i])
	}
	return word
}

func main() {
	fmt.Println(FirstWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Println(FirstWord("  lorem,ipsum  "))
	fmt.Println(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Println(LastWord("  lorem,ipsum  "))
	fmt.Println(LastWord("Write a program that takes a string and displays its last word, followed by a newline     "))
}
