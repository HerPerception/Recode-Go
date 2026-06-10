package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run map.go <input string>")
		return
	}
	str := strings.ToLower(os.Args[1])
	if strings.HasSuffix(os.Args[1], ".txt") {
		data, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		str = strings.ToLower(string(data))
	} //else {
	// 	str = strings.ToLower(os.Args[1])
	// }

	charMap := make(map[rune]int)
	slice := []rune{}
	for _, ch := range str {
		if ch == ' ' {
			continue
		}
		if val, exists := charMap[ch]; exists {
			charMap[ch] = val + 1
		} else {
			charMap[ch] = 1
			slice = append(slice, ch)
		}

		// fmt.Printf("%c: %d\n", ch, charMap[ch])

	}
	for _, ch := range slice {
		fmt.Printf("%c -> %d\n", ch, charMap[ch])
	}
}
