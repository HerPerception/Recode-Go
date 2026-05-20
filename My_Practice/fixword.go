package main

import (
	"fmt"
	"unicode"
)

func fixWord(s string) string {
	word := ""
	for index, ch := range s {
		if index%2 == 0 && unicode.IsLetter(ch) {
			ch -= 32
		}
		word += string(ch)
	}
	return word
}
func main() {
	fmt.Println(fixWord("country"))
	fmt.Println(fixWords("hello"))
}

func fixWords(s string) string {
	word := ""
	for index, ch := range s {
		if index == 0 && unicode.IsLetter(ch) {
			ch -= 32
		} else if index == len(s)-1 && unicode.IsLetter(ch) {
			ch -= 32
		}
		word += string(ch)
	}
	return word
}

//country ===> CoUnTrY

//hello ===> HellO
