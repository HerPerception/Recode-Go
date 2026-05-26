package main

import (
	"fmt"
	"os"
)

func FirstWord() {
	if len(os.Args) != 2 {
		return
	}
	s := os.Args[1]
	word := ""
	for index, ch := range s {
		if ch == ' ' {
			break
		} else if index >= 0 && ch != ' ' {
			word += string(ch)
		}
	}
	fmt.Println(word)
}
