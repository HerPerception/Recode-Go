package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run map.go <input string>")
		return
	}
	str := os.Args[1]
	charMap := make(map[rune]int)
	slice := []rune{}
	for _, ch := range str {
		if val, exists := charMap[ch]; exists {
			charMap[ch] = val + 1
		} else {
			charMap[ch] = 1
			slice = append(slice, ch)
		}

		// fmt.Printf("%c: %d\n", ch, charMap[ch])

	}
	for _, ch := range slice {
		fmt.Printf("%c: %d\n", ch, charMap[ch])
	}
}
