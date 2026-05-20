package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Incorrect number of arguments. Check input and try again.")
		return
	}
	str := os.Args[1]
	first := ""
	last := ""
	count := len(str)
	letters := ""
	digits := ""
	for i := range str {
		if str[i] >= 'A' && str[i] <= 'Z' || str[i] >= 'a' && str[i] <= 'z' {
			letters += string(str[i])
		} else if str[i] >= '0' && str[i] <= '9' {
			digits += string(str[i])
		}
		if i == 0 {
			first = string(str[i])
		} else if i == len(str)-1 {
			last = string(str[i])
		}
	}
	fmt.Printf("Length: %d\n", count)
	fmt.Printf("Letters: %s\n", letters)
	fmt.Printf("Digits: %s\n", digits)
	fmt.Printf("First character: %s\n", first)
	fmt.Printf("Last character: %s\n", last)
}
