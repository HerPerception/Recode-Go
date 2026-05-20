package main

import "fmt"

func countChar(s string, c rune) int {
	count := 0
	for _, ch := range s{
		if ch == c {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countChar("Celebration", 'C'))
}