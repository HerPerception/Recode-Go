package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	word := os.Args[1]
	if len(word) == 0 || word == " " {
		fmt.Println()
		return
	}
	newWord := ""
	str := strings.Fields(word)
	for i := 1; i < len(str); i++ {
		newWord += str[i] + " "
	}
	newWord += str[0]
	fmt.Println(newWord)
}
