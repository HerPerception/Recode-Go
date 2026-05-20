package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

func HashCode(dec string) string {
	word := ""
	for _, ch := range dec {
		ch = rune((int(ch) + len(dec)) % 127)
		if ch < 33 {
			ch += 33
		}
		word += string(ch)
	}
	return word
}
