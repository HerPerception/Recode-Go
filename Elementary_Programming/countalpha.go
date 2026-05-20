package main

import "fmt"

func countAlpha(s string) int {
	count := 0
	for _, ch := range s {
		if ch >= 'A' && ch <= 'z' {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countAlpha("Hello, You ARre welcome 123"))
}