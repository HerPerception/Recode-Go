package main

import (
	"fmt"
)

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}

func Join(str []string, sep string) string {
	word := ""
	for i, ch := range str {
		if i < len(str)-1 {
			word += ch + ":"
		} else {
			word += ch
		}
	}
	return word
}
